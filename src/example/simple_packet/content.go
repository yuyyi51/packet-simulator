package simple_packet

import (
	"fmt"
	"strings"
)

type Content interface {
	Length() int64
	NeedRetransmit() bool
	String() string
}

const SimplePacketDataOverhead = SimplePacketOverHead + 8 // seq + type + size

type SimplePacketContentData struct {
	Offset int64
	Size   int64
}

func (s *SimplePacketContentData) Length() int64 {
	// [Offset 8byte][Data...]
	return 8 + s.Size
}

func (s *SimplePacketContentData) NeedRetransmit() bool {
	return true
}

func (s *SimplePacketContentData) String() string {
	return fmt.Sprintf("Offset: %d, Size: %d", s.Offset, s.Size)
}

type AckRange struct {
	left, right int64
}

func (r *AckRange) String() string {
	return fmt.Sprintf("[%d, %d]", r.left, r.right)
}

type SimplePacketContentAck struct {
	Ranges []AckRange
}

func (s *SimplePacketContentAck) Length() int64 {
	// [left 8byte][right 8byte][left 8byte][right 8byte]...
	return 16 * int64(len(s.Ranges))
}

func (s *SimplePacketContentAck) NeedRetransmit() bool {
	return false
}

func (s *SimplePacketContentAck) String() string {
	sb := strings.Builder{}
	for i, r := range s.Ranges {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(r.String())
	}
	return sb.String()
}

type SimplePacketContentPing struct {
}

func (s *SimplePacketContentPing) Length() int64 {
	return 0
}

func (s *SimplePacketContentPing) NeedRetransmit() bool {
	return false
}

func (s *SimplePacketContentPing) String() string {
	return ""
}

type SimplePacketContentClose struct {
}

func (s *SimplePacketContentClose) Length() int64 {
	return 0
}

func (s *SimplePacketContentClose) NeedRetransmit() bool {
	return true
}

func (s *SimplePacketContentClose) String() string {
	return ""
}
