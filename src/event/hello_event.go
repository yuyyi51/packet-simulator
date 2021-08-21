package event

import (
	"time"

	"github.com/yuyyi51/packet-simulator/src/core"
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
	event.GetOwner().GetLogger().Infof("Hello triggered, helloTime %d, timeOffset %s", event.helloTime, event.GetOwner().GetCurrentTimeOffset())
	if event.helloTime == 0 {
		event.GetOwner().GetLogger().Infof("Hello exit!")
		return
	}
	newEvent := event.GetOwner().CreateEvent(event.GetTriggerTime().Add(time.Second))
	newHelloEvent := &HelloEvent{
		EventI:    newEvent,
		helloTime: event.helloTime - 1,
	}
	err := event.GetOwner().AddEvent(newHelloEvent)
	if err != nil {
		event.GetOwner().GetLogger().Errorf("HelloEvent add new event fail %v", err)
	}
}
