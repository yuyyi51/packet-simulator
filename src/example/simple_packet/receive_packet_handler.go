package simple_packet

type ReceivePacketHandler struct {
	packets        map[int64]struct{}
	smallestTrack  int64
	largestReceive int64
	needAck        bool
}

func NewReceivePacketHandler() *ReceivePacketHandler {
	return &ReceivePacketHandler{
		packets:        make(map[int64]struct{}),
		smallestTrack:  0,
		largestReceive: 0,
	}
}

func (h *ReceivePacketHandler) onPacketReceive(pkt *SimplePacket) {
	if h.smallestTrack > pkt.Seq {
		return
	}
	h.packets[pkt.Seq] = struct{}{}
	if h.largestReceive < pkt.Seq {
		h.largestReceive = pkt.Seq
	}
	if pkt.AckEliciting() {
		h.needAck = true
	}
}

func (h *ReceivePacketHandler) updateSmallestTrack(seq int64) {
	if h.smallestTrack < seq+1 {
		for i := h.smallestTrack; i <= seq; i++ {
			delete(h.packets, i)
		}
		h.smallestTrack = seq + 1
	}
}

func (h *ReceivePacketHandler) generateAckRanges() (ret []AckRange) {
	var left int64
	var lackLeft = true
	for i := h.smallestTrack; i <= h.largestReceive; i++ {
		if _, ok := h.packets[i]; ok {
			if lackLeft {
				left = i
				lackLeft = false
			}
		} else {
			if !lackLeft {
				// it's a range
				right := i - 1
				ret = append(ret, AckRange{left: left, right: right})
				lackLeft = true
			}
		}
	}
	if !lackLeft {
		ret = append(ret, AckRange{left: left, right: h.largestReceive})
	}
	return
}

func (h *ReceivePacketHandler) needSendAck() bool {
	return h.needAck
}
