package core

import (
	"errors"
	"time"
)

type EventManagerI interface {
	CreateEvent(time.Time) EventI
	AddEvent(EventI) error
	Run()
}

type EventManager struct {
	queue       EventQueueI
	currentTime time.Time
	startTime   time.Time
}

func NewEventManager(startTime time.Time) *EventManager {
	return &EventManager{
		queue:       NewEventQueue(),
		currentTime: time.Time{},
	}
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
		newTime := m.waitUntil(event.GetTriggerTime())
		m.setCurrentTime(newTime)
		event.Trigger()
	}
}

func (m *EventManager) setCurrentTime(t time.Time) {
	m.currentTime = t
}

func (m *EventManager) GetCurrentTimeOffset() time.Duration {
	return m.currentTime.Sub(m.startTime)
}

func (m *EventManager) waitUntil(t time.Time) time.Time {
	timer := time.NewTimer(t.Sub(time.Now()))
	newTime := <-timer.C
	return newTime
}
