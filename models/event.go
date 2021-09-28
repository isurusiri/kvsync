package models

// Topic is an event Topic
type Topic string

// Event represents an event
// object exposed by the api.
type Event struct {
	Topic      Topic
	Type       string
	Key        string
	Namespace  string
	FilterKeys []string
	Index      uint64
	Payload    map[string]interface{}
}
