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

	/*
			byteSlice, _ := json.Marshal(order)
		    log.Info(string(byteSlice))

		    newOrder := &Order{}
		    err := json.Unmarshal(byteSlice, &newOrder)
		    if err != nil {
		        log.Fatal(err)
			}

			for _, event := range newOrder.Changes {
		        log.Info("event: ", event)
			}
	*/
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

	order.Status = deserializedOrder.Status
	order.Version = deserializedOrder.Version
	order.Changes = deserializedOrder.Changes

	events := db.GetEventsForAggregate(deserializedOrder.UUID, deserializedOrder.Version)

	log.Info("deserializedOrder: ", deserializedOrder.UUID)
	log.Info("order: ", order.UUID)

	return events

	/*
			for _, event := range deserializedOrder.Changes {
		        log.Info("event: ", event)
			}
	*/
}

func createOrderWithSnapshot(capacity enum.Capacity, official string) *Order {
	order := &Order{}
	db := store.EventStore{}

	//Todo: this is only for test purposes! Needs to be done in aggregate!
	//createOrder and confirmOrder should not return anything
	event := order.createOrder(capacity)
	log.Info("serialized event: ", event)

	db.CreateEvent(1, order.UUID, "order.created", "order", marshalToJSON(event), time.Now())

	MarkOrderAsCommitted(order)

	//so, err := json.Marshal(order)
	//if(err!=nil) {
	//	log.Fatal("Error serializing aggregate: ", err)
	//}
	db.CreateSnapshot(order.UUID, string(marshalToJSON(order)), order.Version)

	event = order.confirmOrder(official)
	db.CreateEvent(2, order.UUID, "order.confirmed", "order", marshalToJSON(event), time.Now())

	MarkOrderAsCommitted(order)

	return order
}

func fulfillOrder(capacity enum.Capacity, official string) *Order {
	order := &Order{}
	db := store.EventStore{}

	log.Info("Creating new Order...")

	event := order.createOrder(capacity)

	db.CreateEvent(1, order.UUID, "order.created", "order", marshalToJSON(event), time.Now())

	event = order.confirmOrder(official)

	db.CreateEvent(2, order.UUID, "order.confirmed", "order", marshalToJSON(event), time.Now())

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
