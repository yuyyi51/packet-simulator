package main

import (
	"github.com/yuyyi51/packet-simulator/src/example/simple_packet"
	"github.com/yuyyi51/packet-simulator/src/simulate"
	"github.com/yuyyi51/packet-simulator/src/utils"
)

func main() {
	manager := simulate.NewSimulateManager()
	d1 := manager.CreateDevice()
	d1.SetAddr("10.0.0.0")

	d2 := manager.CreateDevice()
	d2.SetAddr("10.0.0.1")

	r1 := manager.CreateRouter(1024*1024, 1000)
	r1.SetAddr("10.1.0.0")

	l1 := manager.ConnectDevice(d1, r1)
	l1.SetLoss(utils.NewFraction(10, 100))
	l2 := manager.ConnectDevice(d2, r1)

	_ = l1
	_ = l2

	ab1 := manager.CreateApplicationBase(1440)
	app1 := simple_packet.NewApplication(ab1, 0xfabcdeff, simulate.NewAddr("10.0.0.1", 1440))
	d1.RunApplication(app1)

	ab2 := manager.CreateApplicationBase(1440)
	app2 := simple_packet.NewApplication(ab2, 0xffffffff, simulate.NewAddr("10.0.0.0", 1440))
	d2.RunApplication(app2)

	//app1.Start(simulate.NewAddr("10.0.0.1", 1440))
	app1.Start()
	app1.Write(40 * 1024)
	app2.Start()
	manager.Run()
}
