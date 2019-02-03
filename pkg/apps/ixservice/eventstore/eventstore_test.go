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

	orderCreatedPayload := &event.OrderCreated {
		Capacity: enum.ThreeGB,
	}

	orderConfirmedPayload := &event.OrderConfirmed {
		ConfirmedBy: "me",
	}

	orderRevisedPayload := &event.OrderRevised {
		RevisedBy: "me",
		Reason:    "because I can",
	}

	event.RegisterEvent(orderCreatedPayload)
	event.RegisterEvent(orderConfirmedPayload)
	event.RegisterEvent(orderRevisedPayload)	

	//simulate the creation of a new order aggregate by just generating a new UUID
	aggregateUUID := uuid.NewV4().String()

	createdEvent := event.BuildEvent(orderCreatedPayload, aggregateUUID)
	confirmedEvent := event.BuildEvent(orderConfirmedPayload, aggregateUUID)
	revisedEvent := event.BuildEvent(orderRevisedPayload, aggregateUUID)

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