package models

// EventStream represents the contract
// exposed by the events steaming api.
type EventStream struct {
	Index  uint64
	Events []Event
	Err    error
}
