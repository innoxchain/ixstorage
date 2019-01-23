package main

import (
	"encoding/json"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event"
	store "github.com/innoxchain/ixstorage/pkg/apps/ixservice/eventstore"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	/*
		history := []event.DomainEvent{
			&event.OrderCreatedEvent{AggregateID: "12345", CreatedAt: time.Now().String(), Capacity: enum.SixGB},
			&event.OrderConfirmedEvent{AggregateID: "12345", CreatedAt: time.Now().String(), ConfirmedBy: "me"},
		}

		order := GetOrderFromHistory(history)
		log.Info("Order Aggregate from history: ", order)
	*/

	//fulfillOrder(enum.TenGB, "me")
	//fulfillOrder(enum.ThreeGB, "me too")

	newOrder := createOrderWithSnapshot(enum.SixGB, "me again")
	eventHistory := recreateFromSnapshot(newOrder)
	GetOrderFromHistory(eventHistory, newOrder)

	log.Info("Order Aggregate from eventhistory: ", newOrder)
}

func recreateFromSnapshot(order *Order) []event.DomainEvent {
	db := store.EventStore{}

	s := db.GetSnapshot(order.UUID)

	log.Info("Snapshot serialized: ", s)

	deserializedOrder := &Order{}
	err := json.Unmarshal([]byte(s), deserializedOrder)
	if err != nil {
		log.Fatal("Error deserializing aggregate! ", err)
	}

	log.Info("Deserialized Order: ", deserializedOrder)

	//Todo: order attributes must not be overridden when events have not yet been
	//stored in database

	events := make([]event.DomainEvent, 0)

	if(EventsPublished(order)) {
		order.Status = deserializedOrder.Status
		order.Version = deserializedOrder.Version
		order.Changes = deserializedOrder.Changes
		
		events = db.GetEventsForAggregate(deserializedOrder.UUID, deserializedOrder.Version)
	}

	return events
}

func createOrderWithSnapshot(capacity enum.Capacity, official string) *Order {
	order := &Order{}
	db := store.EventStore{}

	//Todo: this is only for test purposes! Needs to be done in aggregate!
	//createOrder and confirmOrder should not return anything
	event := order.createOrder(capacity)
	log.Info("serialized event: ", event)

	db.CreateEvent(1, order.UUID, "OrderCreated", "order", marshalToJSON(event), time.Now())

	MarkOrderAsCommitted(order)

	db.CreateSnapshot(order.UUID, string(marshalToJSON(order)), order.Version)

	event = order.confirmOrder(official)
	db.CreateEvent(2, order.UUID, "OrderConfirmed", "order", marshalToJSON(event), time.Now())

	MarkOrderAsCommitted(order)

	return order
}

func fulfillOrder(capacity enum.Capacity, official string) *Order {
	order := &Order{}
	db := store.EventStore{}

	log.Info("Creating new Order...")

	event := order.createOrder(capacity)

	db.CreateEvent(1, order.UUID, "OrderCreated", "order", marshalToJSON(event), time.Now())

	event = order.confirmOrder(official)

	db.CreateEvent(2, order.UUID, "OrderConfirmed", "order", marshalToJSON(event), time.Now())

	MarkOrderAsCommitted(order)

	log.Info("Order created: ", order)

	return order
}

func logEvents() {
	db := store.EventStore{}

	log.Info("Reading Events from Eventstore...")
	rs := db.GetEvents()

	for _, it := range rs {
		log.Info(it)
	}
}

func marshalToJSON(object interface{}) string {
	res, err := json.Marshal(object)
	if err != nil {
		log.Fatal("Error serializing aggregate: ", err)
	}
	return string(res)
}
