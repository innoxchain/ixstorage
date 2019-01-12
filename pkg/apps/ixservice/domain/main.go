package main

import (
	"encoding/json"
	store "github.com/innoxchain/ixstorage/pkg/apps/ixservice/eventstore"
	log "github.com/sirupsen/logrus"
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
	log.Info("Order Aggregate from history: ", order)

	order.createOrder(enum.ThreeGB)
	log.Info("Order Aggregate after creating new order", order)

	db := store.EventStore{}
	
	log.Info("Creating new Events...")

	order.createOrder(enum.TenGB)
	order.confirmOrder("me")
	db.CreateEvent(3, order.UUID, "order.created", "order", "{\"test3\":\"test3\"}", time.Now())
	db.CreateEvent(4, order.UUID, "order.confirmed", "order", "{\"test4\":\"test4\"}", time.Now())
	
	log.Info("Order Aggregate, several events later...", order)

	log.Info("Reading Events from Eventstore...")
	rs := db.GetEvents()

	for _, it := range rs {
		log.Info(it)
	}
}

func createOrderWithSnapshot(order *Order, db *store.EventStore) {
	order.createOrder(enum.SixGB)
	db.CreateEvent(1, order.UUID, "order.created", "order", "{\"test\":\"test\"}", time.Now())
	
	so, err := json.Marshal(order)
	if(err!=nil) {
		log.Info("Error serializing aggregate: ", order)
	}
	db.CreateSnapshot(order.UUID, string(so), order.Version)

	order.confirmOrder("me")
	db.CreateEvent(2, order.UUID, "order.confirmed", "order", "{\"test2\":\"test2\"}", time.Now())
}