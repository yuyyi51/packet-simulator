package simulate

import "net"

type PacketI interface {
	GetLength() int64
	GetTargetAddr() net.Addr
	GetSourceAddr() net.Addr
}

type PacketType int

const (
	PacketTypePayload = iota
	PacketTypeAck
)

type Packet struct {
	packetType PacketType
	seq        int64
	length     int64
	sourceAddr net.Addr
	targetAddr net.Addr
}

func (pkt *Packet) GetLength() int64 {
	return pkt.length
}

func (pkt *Packet) GetTargetAddr() net.Addr {
	return pkt.targetAddr
}

func (pkt *Packet) GetSourceAddr() net.Addr {
	return pkt.sourceAddr
}
