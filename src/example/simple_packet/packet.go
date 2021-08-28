package simple_packet

import (
	"fmt"
	"time"

	"github.com/yuyyi51/packet-simulator/src/simulate"
)

type SimplePacketType int

const (
	SimplePacketTypeData SimplePacketType = iota
	SimplePacketTypeAck
	SimplePacketTypePing
	SimplePacketTypeClose
)

func (t SimplePacketType) String() string {
	switch t {
	case SimplePacketTypeData:
		return "Data"
	case SimplePacketTypeAck:
		return "Ack"
	case SimplePacketTypePing:
		return "Ping"
	case SimplePacketTypeClose:
		return "Close"
	}
	return "Undefined"
}

const SimplePacketOverHead = 20 // conv + seq + type

type SimplePacket struct {
	*simulate.PacketBase
	Conv       int64
	Seq        int64
	PacketType SimplePacketType
	Content    Content

	sentTime time.Time
}

func (pkt *SimplePacket) GetLength() int64 {
	// [Conv 8byte][Seq 8byte][Type 4byte][Content...]
	return SimplePacketOverHead + pkt.Content.Length()
}

func (pkt *SimplePacket) String() string {
	return fmt.Sprintf("Packet [%d], type %s, %s", pkt.Seq, pkt.PacketType, pkt.Content)
}

func (pkt *SimplePacket) AckEliciting() bool {
	switch pkt.PacketType {
	case SimplePacketTypeData:
		return true
	case SimplePacketTypePing:
		return true
	case SimplePacketTypeClose:
		return true
	}
	return false
}
