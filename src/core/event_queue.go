package core

type EventQueueI interface {
	AddEvent(EventI)
	PeekEvent() EventI
	PopEvent() EventI
}

type EventQueue struct {
	eventList []EventI
}

func NewEventQueue() *EventQueue {
	return &EventQueue{eventList: make([]EventI, 0)}
}

func (queue *EventQueue) sortEventList() {
	for i := 0; i < len(queue.eventList); i++ {
		for j := 0; j < i; j++ {
			if queue.eventList[i].GetTriggerTime().Before(queue.eventList[j].GetTriggerTime()) {
				t := queue.eventList[i]
				queue.eventList[i] = queue.eventList[j]
				queue.eventList[j] = t
			}
		}
	}
}

func (queue *EventQueue) AddEvent(event EventI) {
	queue.eventList = append(queue.eventList, event)
	queue.sortEventList()
}

func (queue *EventQueue) PeekEvent() EventI {
	if len(queue.eventList) == 0 {
		return nil
	}
	return queue.eventList[0]
}

func (queue *EventQueue) PopEvent() EventI {
	if len(queue.eventList) == 0 {
		return nil
	}
	e := queue.eventList[0]
	queue.eventList = queue.eventList[1:]
	return e
}
