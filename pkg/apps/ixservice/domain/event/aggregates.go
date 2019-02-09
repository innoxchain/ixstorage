package event

import (
	"time"
	//"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event"
)

type Aggregate interface {
	GetAggregateID() string
	trackChanges(e DomainEvent)
	incrementVersion()
}

type BaseAggregate struct {
	UUID    string              `json:"uuid"`
	Version int                 `json:"version"`
	LastModified time.Time		`json:"lastModified"`
	Changes []DomainEvent 		`json:"changes"`
}

func (a *BaseAggregate) GetAggregateID() string {
	return a.UUID
}

func (a *BaseAggregate) incrementVersion() {
	a.Version+=1
}

func (a *BaseAggregate) trackChanges(e Event) {
	a.LastModified=e.GetCreatedAt()
	a.Changes=append(a.Changes, e.Payload.(DomainEvent))
}