package event

//DomainEvent is the domain interface for any kind of events used between client and server communication
type DomainEvent interface {
	GetType() string
	GetAggregateID() string
	GetCreatedAt() string
}
