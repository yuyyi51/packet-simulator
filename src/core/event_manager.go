package core

import (
	"errors"
	"time"

	"github.com/yuyyi51/packet-simulator/src/log"
)

type EventManagerI interface {
	CreateEvent(time.Time) EventI
	AddEvent(EventI) error
	Run()
	GetCurrentTime() time.Time
	GetCurrentTimeOffset() time.Duration
}

type TimeWaiter interface {
	waitUntil(t time.Time) time.Time
}

type EventManager struct {
	queue       EventQueueI
	waiter      TimeWaiter
	currentTime time.Time
	startTime   time.Time
}

func NewEventManager(startTime time.Time) *EventManager {
	m := &EventManager{
		queue:       NewEventQueue(),
		currentTime: startTime,
		startTime:   startTime,
	}
	m.waiter = m
	return m
}

func (m *EventManager) CreateEvent(triggerTime time.Time) EventI {
	event := &BaseEvent{owner: m, triggerTime: triggerTime}
	return event
}

func (m *EventManager) AddEvent(event EventI) error {
	if event.GetOwner() != m {
		return errors.New("event not belong this EventManager")
	}
	m.queue.AddEvent(event)
	return nil
}

func (m *EventManager) Run() {
	for m.queue.PeekEvent() != nil {
		event := m.queue.PopEvent()
		newTime := m.waiter.waitUntil(event.GetTriggerTime())
		m.setCurrentTime(newTime)
		log.Debugf("EventManager running event %#v", event)
		event.Trigger()
	}
}

func (m *EventManager) setCurrentTime(t time.Time) {
	m.currentTime = t
}

func (m *EventManager) GetCurrentTime() time.Time {
	return m.currentTime
}

func (m *EventManager) GetCurrentTimeOffset() time.Duration {
	return m.currentTime.Sub(m.startTime)
}

func (m *EventManager) SetWaiter(w TimeWaiter) {
	m.waiter = w
}

func (m *EventManager) waitUntil(t time.Time) time.Time {
	return t
}
