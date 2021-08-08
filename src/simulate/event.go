package simulate

import "github.com/yuyyi51/packet-simulator/src/core"

type PacketTransportEvent struct {
	core.EventI
	target DeviceI
	pkt    PacketI
}

func NewPacketTransportEvent(e core.EventI, target DeviceI, pkt PacketI) *PacketTransportEvent {
	return &PacketTransportEvent{
		EventI: e,
		target: target,
		pkt:    pkt,
	}
}

func (event *PacketTransportEvent) Trigger() {
	event.target.ReceivePacket(event.pkt)
}

type PacketProcessEvent struct {
	core.EventI

	router *Router
	pkt    PacketI
}

func NewPacketProcessEvent(e core.EventI, router *Router, pkt PacketI) *PacketProcessEvent {
	return &PacketProcessEvent{
		EventI: e,
		router: router,
		pkt:    pkt,
	}
}

func (event *PacketProcessEvent) Trigger() {
	event.router.SendPacket(event.pkt)
	event.router.processNextPacket()
}
