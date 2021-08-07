package core

import "time"

type EventI interface {
	GetTriggerTime() time.Time
	GetOwner() EventManagerI
	Trigger()
}

type BaseEvent struct {
	triggerTime time.Time
	owner       EventManagerI
}

func (e *BaseEvent) GetOwner() EventManagerI {
	return e.owner
}

func (e *BaseEvent) GetTriggerTime() time.Time {
	return e.triggerTime
}

func (e *BaseEvent) Trigger() {
	panic("BaseEvent should never be triggered")
}
