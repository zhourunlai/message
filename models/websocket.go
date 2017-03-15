package models

import (
	"container/list"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct {
	Type      EventType // JOIN, LEAVE, MESSAGE
	User      string
	Timestamp int
	Content   string
}

const wsSize = 20

// Event ws.
var ws = list.New()

// NewWs saves new event to ws list.
func NewWs(event Event) {
	if ws.Len() >= wsSize {
		ws.Remove(ws.Front())
	}
	ws.PushBack(event)
}

// GetEvents returns all events after lastReceived.
func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, ws.Len())
	for event := ws.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
