package event

import (
	"time"

	"github.com/yuyyi51/packet-simulator/src/core"
	"github.com/yuyyi51/packet-simulator/src/log"
)

type HelloEvent struct {
	core.EventI
	helloTime int
}

func NewHelloEvent(e core.EventI, helloTime int) *HelloEvent {
	return &HelloEvent{
		EventI:    e,
		helloTime: helloTime,
	}
}

func (event *HelloEvent) Trigger() {
	log.Infof("Hello triggered, helloTime %d, timeOffset %s", event.helloTime, event.GetOwner().GetCurrentTimeOffset())
	if event.helloTime == 0 {
		log.Infof("Hello exit!")
		return
	}
	newEvent := event.GetOwner().CreateEvent(event.GetTriggerTime().Add(time.Second))
	newHelloEvent := &HelloEvent{
		EventI:    newEvent,
		helloTime: event.helloTime - 1,
	}
	err := event.GetOwner().AddEvent(newHelloEvent)
	if err != nil {
		log.Errorf("HelloEvent add new event fail %v", err)
	}
}
