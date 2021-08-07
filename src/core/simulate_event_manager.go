package core

import (
	"time"
)

type SimulateEventManager struct {
	EventManager
}

func NewSimulateEventManager(mockStartTime time.Time) *SimulateEventManager {
	m := &SimulateEventManager{
		*NewEventManager(mockStartTime),
	}
	m.SetWaiter(m)
	return m
}

func (m *SimulateEventManager) waitUntil(t time.Time) time.Time {
	return t
}
