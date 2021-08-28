package simple_packet

import "time"

type SentPacketHandler struct {
	packets     map[int64]*SimplePacket
	nextSeq     int64
	lostPackets []*SimplePacket
	rttStats    *RttStats
}

func NewSentPacketHandler() *SentPacketHandler {
	return &SentPacketHandler{
		packets:  make(map[int64]*SimplePacket),
		nextSeq:  1,
		rttStats: NewRttStats(),
	}
}

func (h *SentPacketHandler) getNextSeq() int64 {
	s := h.nextSeq
	h.nextSeq++
	return s
}

func (h *SentPacketHandler) onSentPacket(pkt *SimplePacket, current time.Time) {
	_, ok := h.packets[pkt.Seq]
	if ok {
		return
	}
	pkt.sentTime = current
	h.packets[pkt.Seq] = pkt
}

func (h *SentPacketHandler) onAckPacket(seq int64, current time.Time, needUpdateRtt bool) bool {
	pkt, ok := h.packets[seq]
	if !ok {
		return false
	}
	if needUpdateRtt {
		rtt := current.Sub(pkt.sentTime)
		h.rttStats.NewRtt(rtt)
	}
	delete(h.packets, seq)
	return needUpdateRtt
}

func (h *SentPacketHandler) onLostPacket(seq int64) {
	_, ok := h.packets[seq]
	if !ok {
		return
	}
	pkt := h.packets[seq]
	// judge whether need retransmit
	if pkt.Content.NeedRetransmit() {
		h.lostPackets = append(h.lostPackets, pkt)
	}
	delete(h.packets, seq)
}

func (h *SentPacketHandler) queueLostPacket() *SimplePacket {
	var pkt *SimplePacket
	if len(h.lostPackets) != 0 {
		pkt = h.lostPackets[0]
		h.lostPackets = h.lostPackets[1:]
	}
	return pkt
}
