package simulate

import (
	"net"

	"github.com/yuyyi51/packet-simulator/src/log"
)

type ApplicationBaseI interface {
	SendPacket(pkt PacketI)
	SetSendConn(sender PacketSender)
	Port() int
	Addr() net.Addr
}

type ApplicationI interface {
	ApplicationBaseI
	ReceivePacket(pkt PacketI)
}

type PacketSender interface {
	SendPacket(PacketI)
	Addr() string
}

type ApplicationBase struct {
	manager  SimulateManagerI
	sendConn PacketSender
	port     int
	address  net.Addr
}

func NewApplicationBase(port int) *ApplicationBase {
	return &ApplicationBase{port: port}
}

func (app *ApplicationBase) SendPacket(pkt PacketI) {
	app.sendConn.SendPacket(pkt)
}

func (app *ApplicationBase) SetSendConn(sender PacketSender) {
	app.sendConn = sender
}

func (app *ApplicationBase) Port() int {
	return app.port
}

func (app *ApplicationBase) Addr() net.Addr {
	if app.address == nil {
		app.address = NewAddr(app.sendConn.Addr(), app.port)
	}
	return app.address
}

type HelloApplication struct {
	ApplicationBaseI
	maxHelloTime int
}

func NewHelloApplication(base ApplicationBaseI, maxHelloTime int) *HelloApplication {
	return &HelloApplication{
		ApplicationBaseI: base,
		maxHelloTime:     maxHelloTime,
	}
}

func (app *HelloApplication) ReceivePacket(pkt PacketI) {
	if app.maxHelloTime > 0 {
		log.Debugf("ReceivePacket, maxHelloTime %d", app.maxHelloTime)
		app.maxHelloTime--
		echo := &Packet{length: 1000, targetAddr: pkt.GetSourceAddr(), sourceAddr: app.Addr()}
		app.SendPacket(echo)
		return
	}
}

func (app *HelloApplication) Start(addr net.Addr) {
	echo := &Packet{length: 1000, targetAddr: addr, sourceAddr: app.Addr()}
	app.SendPacket(echo)
}
