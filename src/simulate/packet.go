package simulate

import "net"

type PacketI interface {
	GetLength() int64
	GetTargetAddr() net.Addr
	GetSourceAddr() net.Addr
}

type PacketBase struct {
	SourceAddr net.Addr
	TargetAddr net.Addr
}

func (pkt *PacketBase) GetLength() int64 {
	panic("PacketBase shouldn't call GetLength")
}

func (pkt *PacketBase) GetTargetAddr() net.Addr {
	return pkt.TargetAddr
}

func (pkt *PacketBase) GetSourceAddr() net.Addr {
	return pkt.SourceAddr
}

type HelloPacket struct {
	*PacketBase
	length int64
}

func (pkt *HelloPacket) GetLength() int64 {
	return pkt.length
}
