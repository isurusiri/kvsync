package models

// Job represents a nomad apis job contract.
// WIP..
type Job struct {
	Region           *string
	Namespace        *string
	ID               *string
	Name             *string
	Type             *string
	Priority         *int
	AllAtOnce        *bool
	Datacenters      []string
	Meta             map[string]string
	ConsulToken      *string
	VaultToken       *string
	Stop              *bool
	ParentID          *string
	Dispatched        bool
	Payload           []byte
	ConsulNamespace   *string
	VaultNamespace    *string
	NomadTokenID      *string
	Status            *string
	StatusDescription *string
	Stable            *bool
	Version           *uint64
	SubmitTime        *int64
	CreateIndex       *uint64
	ModifyIndex       *uint64
	JobModifyIndex    *uint64
}