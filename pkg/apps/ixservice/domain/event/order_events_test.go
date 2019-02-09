package event

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"fmt"
	"github.com/satori/go.uuid"
)


func TestDeserializeOrderRevisedEvent(t *testing.T) {
	orderRevisedPayload := &OrderRevised {
		RevisedBy: "me",
		Reason:    "not enough money",
	}

	RegisterEvent(orderRevisedPayload)

	pe := PersistentEvent{
		AggregateID:   uuid.NewV4().String(),
		AggregateType: orderRevisedPayload.GetAggregateType(),
		EventType:     orderRevisedPayload.GetEventType(),
		Sequence:	   orderRevisedPayload.GetSequence(),
		CreatedAt:     orderRevisedPayload.GetCreatedAt(),
		RawData:       "{\"RevisedBy\":\"me\",\"Reason\":\"not enough money\"}",
	}

	event, err := pe.Deserialize()

	if err != nil {
		fmt.Println("couldn't deserialize PersistentEvent: ", pe)
	}

	v := reflect.ValueOf(event.Payload)
	kv := v.MapKeys()
	strct := v.MapIndex(kv[0])

	assert.Equal(t, strct.Interface().(string), "me", "RevisedBy must be \"me\"")

	strct = v.MapIndex(kv[1])
	assert.Equal(t, strct.Interface().(string), "not enough money", "Reason must be \"not enough money\"")
}

func TestSerializeOrderRevisedEvent(t *testing.T) {

	orderRevisedPayload := OrderRevised {
		RevisedBy: "me",
		Reason:    "not enough money",
	}

	event := Event{
		AggregateID:   uuid.NewV4().String(),
		AggregateType: orderRevisedPayload.GetAggregateType(),
		EventType:     orderRevisedPayload.GetEventType(),
		Sequence:	   orderRevisedPayload.GetSequence(),
		CreatedAt:     orderRevisedPayload.GetCreatedAt(),
		Payload:	   orderRevisedPayload,
	}

	persistentEvent, err := event.Serialize()

	if err != nil {
		fmt.Println("couldn't serialize Event: ", event)
	}

	expected := "{\"RevisedBy\":\"me\",\"Reason\":\"not enough money\"}"

	assert.Equal(t, expected, persistentEvent.RawData, "Expected result doesn't match!")
}