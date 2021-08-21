package simulate

import (
	"github.com/yuyyi51/packet-simulator/src/utils"
)

type DeviceI interface {
	ReceivePacket(PacketI)
	SendPacket(PacketI)
	Addr() string
	AddLink(LinkI)
}

type PeerDeviceI interface {
	DeviceI
	RunApplication(app ApplicationI)
}

type Device struct {
	manager SimulateManagerI
	link    LinkI
	address string
	appMap  map[int]ApplicationI
}

func NewDevice() *Device {
	return &Device{
		appMap: make(map[int]ApplicationI),
	}
}

func (d *Device) ReceivePacket(pkt PacketI) {
	port := utils.ParsePort(pkt.GetTargetAddr().String())
	app, ok := d.appMap[port]
	if !ok {
		d.manager.GetLogger().Errorf("Device ReceivePacket not found application, %s, %v", pkt.GetTargetAddr(), d.appMap)
		return
	}
	app.ReceivePacket(pkt)
}

func (d *Device) SendPacket(pkt PacketI) {
	d.link.TransportPacket(d, pkt)
}

func (d *Device) Addr() string {
	return d.address
}

func (d *Device) AddLink(l LinkI) {
	d.link = l
}

func (d *Device) RunApplication(app ApplicationI) {
	app.SetSendConn(d)
	if _, ok := d.appMap[app.Port()]; ok {
		//fixme: error, port already in use
	}
	d.appMap[app.Port()] = app
}

func (d *Device) StopApplication(app ApplicationI) {
	delete(d.appMap, app.Port())
}

func (d *Device) SetAddr(addr string) {
	d.address = addr
}
