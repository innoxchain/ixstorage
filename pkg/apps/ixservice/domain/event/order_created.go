package event

import(
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
)

//OrderCreatedEvent is the event initiated when a new order has been created
type OrderCreatedEvent struct {
	AggregateID   string
	CreatedAt     string
	Capacity 	  enum.Capacity
}

func (e *OrderCreatedEvent) GetType() string {
	return "order.created"
}

func (e *OrderCreatedEvent) GetAggregateID() string {
	return e.AggregateID
}

func (e *OrderCreatedEvent) GetCreatedAt() string {
	return e.CreatedAt
}
