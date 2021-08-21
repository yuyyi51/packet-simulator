package simple_packet

type Content interface {
	Length() int64
	NeedRetransmit() bool
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

type AckRange struct {
	left, right int64
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

type SimplePacketContentPing struct {
}

func (s *SimplePacketContentPing) Length() int64 {
	return 0
}

func (s *SimplePacketContentPing) NeedRetransmit() bool {
	return false
}

type SimplePacketContentClose struct {
}

func (s *SimplePacketContentClose) Length() int64 {
	return 0
}

func (s *SimplePacketContentClose) NeedRetransmit() bool {
	return true
}
