package main

import (
	"time"

	"github.com/yuyyi51/packet-simulator/src/core"

	"github.com/yuyyi51/packet-simulator/src/log"

	"github.com/yuyyi51/packet-simulator/src/event"
)

func main() {
	manager := core.NewSimulateEventManager(time.Now())
	e := manager.CreateEvent(time.Now().Add(time.Second * 10))
	helloEvent := event.NewHelloEvent(e, 10)
	err := manager.AddEvent(helloEvent)
	if err != nil {
		log.Fatalf("AddEvent fail, err %v", err)
	}
	manager.Run()
}
