package simple_packet

import (
	"net"

	"github.com/yuyyi51/packet-simulator/src/simulate"
)

type Application struct {
	simulate.ApplicationBaseI
	sess *Session
}

func NewApplication(base simulate.ApplicationBaseI, conv int64, remoteAddr net.Addr) *Application {
	app := &Application{ApplicationBaseI: base}
	sess := NewSession(app, conv, remoteAddr)
	app.sess = sess
	return app
}

func (app *Application) ReceivePacket(p simulate.PacketI) {
	pkt := p.(*SimplePacket)
	app.onReceivePacket(pkt)
}

func (app *Application) onReceivePacket(pkt *SimplePacket) {
	app.sess.onPacketReceive(pkt)
}

func (app *Application) Write(length int64) {
	app.sess.SimulateWrite(length)
}

func (app *Application) Start() {
	app.sess.Start()
}
