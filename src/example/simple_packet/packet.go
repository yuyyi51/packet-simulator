package simple_packet

import "github.com/yuyyi51/packet-simulator/src/simulate"

type SimplePacketType int

const (
	SimplePacketTypeData SimplePacketType = iota
	SimplePacketTypeAck
	SimplePacketTypePing
	SimplePacketTypeClose
)

const SimplePacketOverHead = 20 // conv + seq + type

type SimplePacket struct {
	*simulate.PacketBase
	Conv       int64
	Seq        int64
	PacketType SimplePacketType
	Content    Content
}

func (pkt *SimplePacket) GetLength() int64 {
	// [Conv 8byte][Seq 8byte][Type 4byte][Content...]
	return SimplePacketOverHead + pkt.Content.Length()
}
