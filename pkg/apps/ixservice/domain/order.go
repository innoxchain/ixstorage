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
	UUID    string `json:"uuid"`
	Version int `json:"version"`
	Status  enum.OrderStatus `json:"status"`
	Changes []event.DomainEvent `json:"changes"`
}

func (o *Order) createOrder(cap enum.Capacity) {
	o.Version=0
	o.trackChange(&event.OrderCreatedEvent{AggregateID:uuid.NewV4().String(), CreatedAt: time.Now().String(), Capacity:cap})
}

func (o *Order) confirmOrder(issuer string) {
	o.trackChange(&event.OrderConfirmedEvent{AggregateID: o.UUID, CreatedAt: time.Now().String(), ConfirmedBy: issuer})
}

func (o *Order) trackChange(event event.DomainEvent) {
	o.Changes = append(o.Changes, event)
	o.transition(event)
}

func (o *Order) transition(ev event.DomainEvent) {
	switch e := ev.(type) {
	case *event.OrderCreatedEvent:
		o.UUID = e.GetAggregateID()
		o.Status = enum.Created
	case *event.OrderConfirmedEvent:
		o.UUID = e.GetAggregateID()
		o.Status = enum.Confirmed
	}
	o.Version++
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
	(Pending Changes: %d)
	(Version: %d)`

	return fmt.Sprintf(format, o.UUID, o.Status, len(o.Changes), o.Version)
}