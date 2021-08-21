package simulate

import (
	"time"
)

type LinkI interface {
	TransportPacket(source DeviceI, pkt PacketI)
	RemoteAddr(peer DeviceI) string
}

type Link struct {
	manager SimulateManagerI
	peer1   DeviceI
	peer2   DeviceI

	delay time.Duration
}

func NewLink(peer1 DeviceI, peer2 DeviceI) *Link {
	return &Link{
		peer1: peer1,
		peer2: peer2,
	}
}

func (l *Link) RemoteAddr(peer DeviceI) string {
	if peer == l.peer1 {
		return l.peer2.Addr()
	} else if peer == l.peer2 {
		return l.peer1.Addr()
	}
	return ""
}

func (l *Link) TransportPacket(source DeviceI, pkt PacketI) {
	if source == l.peer1 {
		// send packet to peer2
		e := l.manager.CreateEvent(l.manager.GetCurrentTime().Add(l.delay))
		eve := NewPacketTransportEvent(e, l.peer2, pkt)
		_ = l.manager.AddEvent(eve)
		return
	} else if source == l.peer2 {
		// send packet to peer1
		e := l.manager.CreateEvent(l.manager.GetCurrentTime().Add(l.delay))
		eve := NewPacketTransportEvent(e, l.peer1, pkt)
		_ = l.manager.AddEvent(eve)
		return
	}
	// error
	l.manager.GetLogger().Errorf("source not match link, \nsource: %v\n, peer1: %v\n, peer2: %v\n", source, l.peer1, l.peer2)
}
