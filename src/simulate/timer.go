package simulate

import "time"

type Timer struct {
	manager  SimulateManagerI
	interval time.Duration
	cb       func()
	stopped  bool
}

func NewTimer(interval time.Duration, cb func()) *Timer {
	return &Timer{
		interval: interval,
		cb:       cb,
	}
}

func (t *Timer) onFire() {
	if !t.stopped {
		t.cb()
	}
}

func (t *Timer) Start() {
	t.stopped = false
	te := t.manager.CreateEvent(t.manager.GetCurrentTime().Add(t.interval))
	event := NewTimerEvent(te, t)
	t.manager.AddEvent(event)
}

func (t *Timer) Stop() {
	t.stopped = true
}
