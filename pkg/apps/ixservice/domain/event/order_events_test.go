package event

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	//"reflect"
	"testing"
)

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
		fmt.Println("couldn't deserialize PersistentEvent: ", pe)
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
		fmt.Println("couldn't serialize Event: ", event)
	}

	expected := "{\"revisedBy\":\"me\",\"reason\":\"not enough money\"}"

	assert.Equal(t, expected, persistentEvent.RawData, "Expected result doesn't match!")
}
