package event

import (
	"time"
	//"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event"
)

type Aggregate interface {
	GetAggregateID() string
	trackChanges(e Event)
	incrementVersion()
}

type BaseAggregate struct {
	UUID    string              `json:"uuid"`
	Version int                 `json:"version"`
	LastModified time.Time		`json:"lastModified"`
	Changes []Event 			`json:"changes"`
}

func (a *BaseAggregate) GetAggregateID() string {
	return a.UUID
}

func (a *BaseAggregate) incrementVersion() {
	a.Version+=1
}

func (a *BaseAggregate) trackChanges(e Event) {
	a.LastModified=e.CreatedAt
	a.Changes=append(a.Changes, e)
}

func (a *BaseAggregate) MarkAsCommited() {
	a.Changes = nil
	a.Changes = make([]Event, 0)
}