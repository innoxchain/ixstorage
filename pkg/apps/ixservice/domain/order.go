package main

import (
	"time"
	"fmt"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"

	"github.com/satori/go.uuid"
)

//Order is the Aggregate which will control the behaviour of state transitions of this Aggregate.
type Order struct {
	uuid    string
	status  event.Status
	changes []event.DomainEvent
}

func (o *Order) createOrder(cap enum.Capacity) {
	o.trackChange(&event.OrderCreatedEvent{AggregateID:uuid.NewV4().String(), CreatedAt: time.Now().String(), Capacity: enum.TenGB})
}

func (o *Order) trackChange(event event.DomainEvent) {
	o.changes = append(o.changes, event)
	o.transition(event)
}

func (o *Order) transition(ev event.DomainEvent) {
	switch e := ev.(type) {
	case *event.OrderCreatedEvent:
		o.uuid = e.GetAggregateID()
		o.status = event.OrderCreated
	case *event.OrderConfirmedEvent:
		o.uuid = e.GetAggregateID()
		o.status = event.OrderConfirmed
	}
}

func GetOrderFromHistory(events []event.DomainEvent) *Order {
	order := &Order{}
	for _, ev := range events {
		order.transition(ev)
	}
	return order
}

func (o *Order) String() string {
	format := `Order:
	uuid: %s
	status: %s
	(Pending Changes: %d)`

	return fmt.Sprintf(format, o.uuid, o.status, len(o.changes))
}