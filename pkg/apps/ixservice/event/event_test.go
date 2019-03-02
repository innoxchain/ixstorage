package event

import (
	"github.com/satori/go.uuid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainEventRegisterEvents(t *testing.T) {
	orderCreated := &OrderCreated{}
	orderConfirmed := &OrderConfirmed{}

	RegisterEvent(orderCreated)
	RegisterEvent(orderConfirmed)

	assert.Equal(t, orderCreated, getEvent("OrderCreated"), "Registered Event should be event.OrderCreatedEvent")
	assert.Equal(t, orderConfirmed, getEvent("OrderConfirmed"), "Registered Event should be event.OrderConfirmedEvent")
}

func TestDeserializeOrderRevisedEvent(t *testing.T) {
	orderRevisedPayload := &OrderRevised{
		RevisedBy: "me",
		Reason:    "not enough money",
	}

	RegisterEvent(orderRevisedPayload)

	pe := PersistentEvent{
		AggregateID:   uuid.NewV4().String(),
		AggregateType: orderRevisedPayload.GetAggregateType(),
		EventType:     orderRevisedPayload.GetEventType(),
		Sequence:      orderRevisedPayload.GetSequence(),
		CreatedAt:     orderRevisedPayload.GetCreatedAt(),
		RawData:       "{\"RevisedBy\":\"me\",\"Reason\":\"not enough money\"}",
	}

	event, err := pe.Deserialize()

	if err != nil {
		t.Fatal("couldn't deserialize PersistentEvent: ", pe)
	}

	t.Log("Event Payload: ", event.Payload)

	assert.Equal(t, event.Payload.(*OrderRevised).RevisedBy, "me", "RevisedBy must be \"me\"")
	assert.Equal(t, event.Payload.(*OrderRevised).Reason, "not enough money", "Reason must be \"not enough money\"")
}


func TestSerializeOrderRevisedEvent(t *testing.T) {

	orderRevisedPayload := OrderRevised{
		RevisedBy: "me",
		Reason:    "not enough money",
	}

	event := Event{
		AggregateID:   uuid.NewV4().String(),
		AggregateType: orderRevisedPayload.GetAggregateType(),
		EventType:     orderRevisedPayload.GetEventType(),
		Sequence:      orderRevisedPayload.GetSequence(),
		CreatedAt:     orderRevisedPayload.GetCreatedAt(),
		Payload:       orderRevisedPayload,
	}

	persistentEvent, err := event.Serialize()

	if err != nil {
		t.Fatal("couldn't serialize Event: ", event)
	}

	expected := "{\"revisedBy\":\"me\",\"reason\":\"not enough money\"}"

	assert.Equal(t, expected, persistentEvent.RawData, "Expected result doesn't match!")
}
