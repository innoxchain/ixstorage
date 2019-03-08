package event

/*
import (
	"encoding/json"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreEvent(t *testing.T) {
	db := EventStore{}

	orderRevisedPayload := OrderRevised{
		RevisedBy: "me",
		Reason:    "not enough money",
	}

	event := BuildEvent(orderRevisedPayload, uuid.NewV4().String())

	t.Log(event)

	err := db.PersistEvent(event)

	assert.True(t, err == nil, "Insertion to db should not fail")
}

func TestGetEventForSnapshot(t *testing.T) {

	db := EventStore{}

	orderCreated := &OrderCreated{
		UUID:     uuid.NewV4().String(),
		Capacity: ThreeGB,
	}

	orderConfirmed := &OrderConfirmed{
		ConfirmedBy: "me",
	}

	orderRevised := &OrderRevised{
		RevisedBy: "me",
		Reason:    "because I can",
	}

	RegisterEvent(orderCreated)
	RegisterEvent(orderConfirmed)
	RegisterEvent(orderRevised)

	//simulate the creation of a new order aggregate by just generating a new UUID
	aggregateUUID := orderCreated.UUID

	createdEvent := BuildEvent(orderCreated, aggregateUUID)
	confirmedEvent := BuildEvent(orderConfirmed, aggregateUUID)
	revisedEvent := BuildEvent(orderRevised, aggregateUUID)

	db.PersistEvent(createdEvent)
	db.PersistEvent(confirmedEvent)
	db.PersistEvent(revisedEvent)

	events := db.EventsForAggregate(aggregateUUID, 1)

	t.Log("size", len(events))

	assert.True(t, len(events) == 2, "There must be two items in resulting events slice")

	for i := 0; i < len(events); i++ {
		t.Log("event[", i, "]=", events[i])
	}
}

func TestEventsWithAggregates(t *testing.T) {
	//db := EventStore{}

	var order Order

	orderCreated := &OrderCreated{
		UUID:     uuid.NewV4().String(),
		Capacity: ThreeGB,
	}

	RegisterEvent(orderCreated)

	e := BuildEvent(orderCreated, orderCreated.UUID)
	e.ApplyChanges(&order)

	t.Log("OrderCreatedEvent: ", e)
	//t.Log("Order: ", order)

	assert.Equal(t, order.UUID, e.AggregateID, "Created order aggregate's id must match OrderCreatedEvent's id")
	assert.Equal(t, order.Version, 1, "Created order's Version must be 1")
	assert.Equal(t, order.Changes[0], e, "Created order's unpersisted changes must match OrderCreatedEvent's payload")

	orderConfirmed := &OrderConfirmed{
		ConfirmedBy: "me",
	}

	RegisterEvent(orderConfirmed)

	e = BuildEvent(orderConfirmed, order.UUID)
	e.ApplyChanges(&order)

	t.Log("OrderConfirmedEvent: ", e)
	t.Log("Order: ", order)

	//err := db.PersistEvent(e)
	//err := db.Persist(&order.BaseAggregate)

	//assert.True(t, err==nil, "Insertion to db should not fail")
	tmpOrder := order
	tmpOrder.MarkAsCommited()

	Replay(&tmpOrder, order.Changes)

	t.Log("tmpOrder: ", tmpOrder)

	assert.True(t, tmpOrder.Capacity == ThreeGB, "Capacity must be ThreeGB after event replay")
	assert.True(t, tmpOrder.ConfirmedBy == "me", "ConfirmedBy must be me after event replay")

	//order.MarkAsCommited()
	//assert.True(t, len(order.Changes)==0, "There must not be any pending changes anymore")
}

func TestSnapshots(t *testing.T) {

	db := EventStore{}

	var order Order

	//first we create a new event OrderCreated, register it with the
	//event registry and apply the changes to the aggregate
	orderCreated := &OrderCreated{
		UUID:     uuid.NewV4().String(),
		Capacity: ThreeGB,
	}

	RegisterEvent(orderCreated)
	e := BuildEvent(orderCreated, orderCreated.UUID)
	e.ApplyChanges(&order)

	//then we persist the aggregate and it's pending changes
	db.Persist(&order.BaseAggregate)

	//now, we want to create a snapshot of the current aggregate state
	//marshal it to JSON and store it in eventstore
	db.CreateSnapshot(order.UUID, marshalToJSON(order), order.Version)

	//Now, additional events will occur over time and we
	//will create and persist them as usual.
	orderConfirmed := &OrderConfirmed{
		ConfirmedBy: "me",
	}
	RegisterEvent(orderConfirmed)
	e = BuildEvent(orderConfirmed, order.UUID)
	e.ApplyChanges(&order)
	db.Persist(&order.BaseAggregate)

	orderRevised := &OrderRevised{
		RevisedBy: "me",
		Reason:    "because I can",
	}
	RegisterEvent(orderRevised)
	e = BuildEvent(orderRevised, order.UUID)
	e.ApplyChanges(&order)
	db.Persist(&order.BaseAggregate)

	//At this point assume we need to recreate a part of our
	//event history for some reason (maybe we want to find some specific
	//events during a certain period of time).
	//That's where we load our previously created snapshot and replay
	//all events which occured in the meantime. When there's a large
	//event history we might work with snapshots at given points in time
	//to save performance when loading large datasets.
	snapshot := db.GetSnapshot(order.UUID)

	deserializedOrder := &Order{}
	err := json.Unmarshal([]byte(snapshot), deserializedOrder)
	if err != nil {
		log.Fatal("Error deserializing aggregate! ", err)
	}

	t.Log("Deserialized Order: ", deserializedOrder)

	assert.True(t, deserializedOrder.ConfirmedBy=="", "ConfirmedBy must be empty after loading snapshot")
	assert.True(t, deserializedOrder.RevisedStat.RevisedBy=="", "RevisedStat.RevisedBy must be empty after loading snapshot")
	assert.True(t, deserializedOrder.RevisedStat.Reason=="", "RevisedStat.Reason must be empty after loading snapshot")

	events := db.EventsForAggregate(deserializedOrder.UUID, deserializedOrder.Version)

	for _, event := range events {
		event.ApplyChanges(deserializedOrder)
	}

	assert.Equal(t, orderConfirmed.ConfirmedBy, deserializedOrder.ConfirmedBy, "ConfirmedBy must match after replaying events for aggregate")
	assert.Equal(t, orderRevised.RevisedBy, deserializedOrder.RevisedStat.RevisedBy, "RevisedStat.RevisedBy must match after loading snapshot")
	assert.Equal(t, orderRevised.Reason, deserializedOrder.RevisedStat.Reason, "RevisedStat.Reason must match after loading snapshot")
}

func marshalToJSON(object interface{}) string {
	res, err := json.Marshal(object)
	if err != nil {
		log.Fatal("Error serializing aggregate: ", err)
	}
	return string(res)
}
*/