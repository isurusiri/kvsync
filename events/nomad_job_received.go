package events

import (
	"github.com/isurusiri/funnel/models"
)

// NomadJobReceivedPayload is the event that is
// going to trigger when jobs are received.

// NomadJobReceivedPayload is the data passed when
// job received event is triggered.
type NomadJobReceivedPayload struct {
	Job models.Job
}

type nomadJobReceived struct {
	handlers []interface{ Handle(NomadJobReceivedPayload) }
}

// Register adds an event handler for this event.
func (n *nomadJobReceived) Register(handler interface{ Handle(NomadJobReceivedPayload) }) {
	n.handlers = append(n.handlers, handler)
}

// Trigger sends out an event with the payload.
func (n *nomadJobReceived) Trigger(payload NomadJobReceivedPayload) {
	for _, handler := range n.handlers {
		go handler.Handle(payload)
	}
}
