package models

import "time"

// Evaluation represents the api response
// of the evaluation in nomad api.
type Evaluation struct {
	ID                   string
	Priority             int
	Type                 string
	TriggeredBy          string
	Namespace            string
	JobID                string
	JobModifyIndex       uint64
	NodeID               string
	NodeModifyIndex      uint64
	DeploymentID         string
	Status               string
	StatusDescription    string
	Wait                 time.Duration
	WaitUntil            time.Time
	NextEval             string
	PreviousEval         string
	BlockedEval          string
	FailedTGAllocs       map[string]*interface{}
	ClassEligibility     map[string]bool
	EscapedComputedClass bool
	QuotaLimitReached    string
	AnnotatePlan         bool
	QueuedAllocations    map[string]int
	SnapshotIndex        uint64
	CreateIndex          uint64
	ModifyIndex          uint64
	CreateTime           int64
	ModifyTime           int64
}
