package handlers

import (
	"fmt"
	"io"

	"github.com/isurusiri/funnel/actions"
	"github.com/isurusiri/funnel/events"
)

func init() {
	// could be a problem
	nomadEventReceivedHandler := nomadEventReceivedHandler{
		EventBody: nil,
	}
	events.NomadEventReceived.Register(nomadEventReceivedHandler)
}

type nomadEventReceivedHandler struct {
	EventBody io.Reader
}

func (n nomadEventReceivedHandler) Handle(payload events.NomadEventReceivedPayload) {
	fmt.Printf("#%+v\n", payload)
	eventID, jobID := actions.GetJobIDFromEventStream(payload.EventBody)
	fmt.Printf("#%d,%s", eventID, jobID)
	fmt.Println("#---")
}
