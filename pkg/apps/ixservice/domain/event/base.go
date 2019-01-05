package event

//Status is showing the possible event states of an aggregate
type Status int

const (
	//OrderCreated is set when a new order has been created
	OrderCreated Status = iota
	//OrderConfirmed is set when an existing order has been confirmed
	OrderConfirmed Status = iota
)

//DomainEvent is the domain interface for any kind of events used between client and server communication
type DomainEvent interface {
	GetType() string
	GetAggregateID() string
	GetCreatedAt() string
}
