package main

import (
	"fmt"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
	"time"
)

func main() {
	history := []event.DomainEvent{
		&event.OrderCreatedEvent{AggregateID: "12345", CreatedAt: time.Now().String(), Capacity: enum.SixGB},
		&event.OrderConfirmedEvent{AggregateID: "12345", CreatedAt: time.Now().String(), ConfirmedBy: "me"},
	}

	order := GetOrderFromHistory(history)
	fmt.Println("Order Aggregate from history\n", order)

	order.createOrder(enum.ThreeGB)
	fmt.Println("Order Aggregate after creating new order\n", order)
}
