package handlers

import "github.com/isurusiri/funnel/events"

func init() {

}

type nomadJobReceivedHandler struct {
	eventID string
	jobID   string
}

func (n nomadJobReceivedHandler) Handle(payload events.NomadEventReceivedPayload) {

}
