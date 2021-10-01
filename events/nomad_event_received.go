package events

import (
	"github.com/isurusiri/funnel/models"
)

// NomadEventReceived is the event that is going to
// trigger when events are received.
var NomadEventReceived nomadEventReceived

// NomadEventReceivedPayload is the data passed when
// a nomad event received event is triggered.
type NomadEventReceivedPayload struct {
	Event models.EventStream
}

type nomadEventReceived struct {
	handlers []interface{ Handle(NomadEventReceivedPayload) }
}

// Register adds an event handler for this event.
func (n *nomadEventReceived) Register(handler interface{ Handle(NomadEventReceivedPayload) }) {
	n.handlers = append(n.handlers, handler)
}

// Trigger sends out an event with the payload.
func (n *nomadEventReceived) Trigger(payload NomadEventReceivedPayload) {
	for _, handler := range n.handlers {
		go handler.Handle(payload)
	}
}
