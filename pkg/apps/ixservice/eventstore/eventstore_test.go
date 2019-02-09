package eventstore

import (
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
	"github.com/stretchr/testify/assert"
	"github.com/satori/go.uuid"
	//"time"
	"testing"

	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event"
)

func TestStoreEvent(t *testing.T) {
	db := EventStore{}

	orderRevisedPayload := event.OrderRevised {
		RevisedBy: "me",
		Reason:    "not enough money",
	}

	event := event.BuildEvent(orderRevisedPayload, uuid.NewV4().String())

	t.Log(event)

	err := db.PersistEvent(event)

	assert.True(t, err==nil, "Insertion to db should not fail")
}


func TestGetEventForSnapshot(t *testing.T) {

	db := EventStore{}

	orderCreated := &event.OrderCreated {
		UUID: uuid.NewV4().String(),
		Capacity: enum.ThreeGB,
	}

	orderConfirmed := &event.OrderConfirmed {
		ConfirmedBy: "me",
	}

	orderRevised := &event.OrderRevised {
		RevisedBy: "me",
		Reason:    "because I can",
	}

	event.RegisterEvent(orderCreated)
	event.RegisterEvent(orderConfirmed)
	event.RegisterEvent(orderRevised)	

	//simulate the creation of a new order aggregate by just generating a new UUID
	aggregateUUID := orderCreated.UUID

	createdEvent := event.BuildEvent(orderCreated, aggregateUUID)
	confirmedEvent := event.BuildEvent(orderConfirmed, aggregateUUID)
	revisedEvent := event.BuildEvent(orderRevised, aggregateUUID)

	db.PersistEvent(createdEvent)
	db.PersistEvent(confirmedEvent)
	db.PersistEvent(revisedEvent)

	events := db.EventsForAggregate(aggregateUUID, 1)

	t.Log("size", len(events))

	assert.True(t, len(events)==2, "There must be two items in resulting events slice")

	for i := 0; i< len(events); i++ {
		t.Log("event[", i, "]=", events[i])
	}
}

func TestEventsWithAggregates(t *testing.T) {
	db := EventStore{}

	var order event.Order

	orderCreated := &event.OrderCreated {
		UUID: uuid.NewV4().String(),
		Capacity: enum.ThreeGB,
	}

	event.RegisterEvent(orderCreated)

	e := event.BuildEvent(orderCreated, orderCreated.UUID)
	e.ApplyChanges(&order)

	t.Log("OrderCreatedEvent: ", e)
	//t.Log("Order: ", order)

	assert.Equal(t, order.UUID, e.AggregateID, "Created order aggregate's id must match OrderCreatedEvent's id")
	assert.Equal(t, order.Version, 1, "Created order's Version must be 1")
	assert.Equal(t, order.Changes[0], e, "Created order's unpersisted changes must match OrderCreatedEvent's payload")


	orderConfirmed := &event.OrderConfirmed {
		ConfirmedBy: "me",
	}

	event.RegisterEvent(orderConfirmed)

	e = event.BuildEvent(orderConfirmed, order.UUID)
	e.ApplyChanges(&order)

	t.Log("OrderConfirmedEvent: ", e)
	//t.Log("Order: ", order)

	//err := db.PersistEvent(e)
	err := db.Persist(&order.BaseAggregate)

	assert.True(t, err==nil, "Insertion to db should not fail")

	//order.MarkAsCommited()
	assert.True(t, len(order.Changes)==0, "There must not be any pending changes anymore")
}