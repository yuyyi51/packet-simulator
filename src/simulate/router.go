package simulate

import (
	"time"

	"github.com/yuyyi51/packet-simulator/src/utils"
)

const precision = 1000000

type Router struct {
	manager        SimulateManagerI
	linkMap        map[string]LinkI
	bandwidth      int64 // Byte/s
	pktQueue       []PacketI
	maxPktQueueLen int
	address        string
	isProcessing   bool
}

func NewRouter(bandwidth int64, maxPktQueueLen int) *Router {
	return &Router{
		bandwidth:      bandwidth,
		maxPktQueueLen: maxPktQueueLen,
		linkMap:        make(map[string]LinkI),
	}
}

func (r *Router) addPacketIntoQueue(pkt PacketI) {
	if len(r.pktQueue) >= r.maxPktQueueLen {
		return
	}
	r.pktQueue = append(r.pktQueue, pkt)
}

func (r *Router) ReceivePacket(pkt PacketI) {
	if r.isProcessing {
		r.addPacketIntoQueue(pkt)
	} else {
		r.processPacket(pkt)
	}
}

func (r *Router) SendPacket(pkt PacketI) {
	addr := utils.ParseAddr(pkt.GetTargetAddr().String())
	l, ok := r.linkMap[addr]
	if !ok {
		r.manager.GetLogger().Errorf("Router send packet not found target addr, %s, %v", pkt.GetTargetAddr().String(), r.linkMap)
		return
	}
	l.TransportPacket(r, pkt)
}

func (r *Router) Addr() string {
	return r.address
}

func (r *Router) AddLink(l LinkI) {
	remoteAddr := l.RemoteAddr(r)
	r.linkMap[remoteAddr] = l
}

func (r *Router) processPacket(pkt PacketI) {
	r.isProcessing = true
	cost := time.Second * precision * time.Duration(pkt.GetLength()) / time.Duration(r.bandwidth) / precision
	e := r.manager.CreateEvent(r.manager.GetCurrentTime().Add(cost))
	event := NewPacketProcessEvent(e, r, pkt)
	_ = r.manager.AddEvent(event)
}

func (r *Router) processNextPacket() {
	if len(r.pktQueue) != 0 {
		pkt := r.pktQueue[0]
		r.pktQueue = r.pktQueue[1:]
		r.processPacket(pkt)
	} else {
		r.isProcessing = false
	}
}

func (r *Router) SetAddr(addr string) {
	r.address = addr
}
