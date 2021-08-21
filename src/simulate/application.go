package simulate

import (
	"fmt"
	"net"
	"time"
)

type ApplicationBaseI interface {
	SendPacket(pkt PacketI)
	SetSendConn(sender PacketSender)
	Port() int
	Addr() net.Addr
	Now() time.Time
	CreateTimer(interval time.Duration, cb func()) *Timer

	GetManager() SimulateManagerI
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

func (app *ApplicationBase) Now() time.Time {
	return app.manager.GetCurrentTime()
}

func (app *ApplicationBase) Addr() net.Addr {
	if app.address == nil {
		app.address = NewAddr(app.sendConn.Addr(), app.port)
	}
	return app.address
}

func (app *ApplicationBase) CreateTimer(interval time.Duration, cb func()) *Timer {
	return app.manager.CreateTimer(interval, cb)
}

func (app *ApplicationBase) GetManager() SimulateManagerI {
	return app.manager
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
		fmt.Printf("%s\n", app)
		fmt.Printf("%s\n", app.GetManager())
		fmt.Printf("%s\n", app.GetManager().GetLogger())
		app.GetManager().GetLogger().Debugf("ReceivePacket, maxHelloTime %d", app.maxHelloTime)
		app.maxHelloTime--
		echo := &HelloPacket{length: 1000, PacketBase: &PacketBase{TargetAddr: pkt.GetSourceAddr(), SourceAddr: app.Addr()}}
		app.SendPacket(echo)
		return
	}
}

func (app *HelloApplication) Start(addr net.Addr) {
	echo := &HelloPacket{length: 1000, PacketBase: &PacketBase{TargetAddr: addr, SourceAddr: app.Addr()}}
	app.SendPacket(echo)
}
