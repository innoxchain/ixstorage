package main

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	
	es "github.com/innoxchain/ixstorage/pkg/apps/ixservice/event"
	log "github.com/sirupsen/logrus"
)

var DB = es.EventStore{}

func createSnapshot(order *es.Order) {
	DB.CreateSnapshot(order.GetAggregateID(), marshalToJSON(order), order.GetVersion())
}

func getSnapshot(order es.Order) *es.Order {
	snapshot := DB.GetSnapshot(order.GetAggregateID())

	deserializedOrder := &es.Order{}
	err := json.Unmarshal([]byte(snapshot), deserializedOrder)
	if err != nil {
		log.Fatal("Error deserializing aggregate! ", err)
	}

	return deserializedOrder
}

func replayEvents(order *es.Order) {
	events := DB.EventsForAggregate(order.GetAggregateID(), order.GetVersion())

	es.Replay(order, events)
}

func marshalToJSON(object interface{}) string {
	res, err := json.Marshal(object)
	if err != nil {
		log.Fatal("Error serializing aggregate: ", err)
	}
	return string(res)
}

func main() {

	var order es.Order

	es.RegisterEvent(&es.OrderCreated{})
	es.RegisterEvent(&es.OrderConfirmed{})
	es.RegisterEvent(&es.OrderRevised{})


	createCommand := es.CreateOrderCommand{
		UUID: uuid.NewV4().String(),
		Capacity: es.ThreeGB,
	}

	confirmCommand := es.ConfirmOrderCommand{
		ConfirmedBy: "myself",
	}

	revisedCommand := es.ReviseOrderCommand {
		RevisedBy: "me again",
		Reason: "invalid",
	}

	es.ApplyCommand(createCommand, &order)
	es.ApplyCommand(confirmCommand, &order)

	createSnapshot(&order)

	es.ApplyCommand(revisedCommand, &order)

	snapshot := getSnapshot(order)

	replayEvents(snapshot)

	log.Info("Snapshot after replay: ", snapshot)
}