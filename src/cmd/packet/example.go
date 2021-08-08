package main

import "github.com/yuyyi51/packet-simulator/src/simulate"

func main() {
	manager := simulate.NewSimulateManager()
	d1 := manager.CreateDevice()
	d1.SetAddr("10.0.0.0")

	d2 := manager.CreateDevice()
	d2.SetAddr("10.0.0.1")

	r1 := manager.CreateRouter(1024*1024, 1000)
	r1.SetAddr("10.1.0.0")

	l1 := manager.ConnectDevice(d1, r1)
	l2 := manager.ConnectDevice(d2, r1)

	_ = l1
	_ = l2

	ab1 := simulate.NewApplicationBase(1440)
	app1 := simulate.NewHelloApplication(ab1, 10)
	d1.RunApplication(app1)

	ab2 := simulate.NewApplicationBase(1440)
	app2 := simulate.NewHelloApplication(ab2, 10)
	d2.RunApplication(app2)

	app1.Start(simulate.NewAddr("10.0.0.1", 1440))

	manager.Run()
}
