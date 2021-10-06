package handlers

import "github.com/isurusiri/funnel/events"

func init() {

}

type nomadEventReceivedHandler struct {
	eventID string
	jobID   string
}

func (n nomadEventReceivedHandler) Handle(payload events.NomadEventReceivedPayload) {

}
