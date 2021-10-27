package actions

import (
	"encoding/json"
	"io"
	"log"

	"github.com/isurusiri/funnel/models"
)

// GetJobIDFromEventStream is used to read nomad job and
// extract job id of the event associated.
func GetJobIDFromEventStream(events io.Reader) (uint64, string) {

	var eventStream models.EventStream
	err := json.NewDecoder(events).Decode(&eventStream)
	if err != nil {
        log.Fatal(err)
    }

	return eventStream.Index, eventStream.Events[0].Payload["Evaluation"].JobID
}
