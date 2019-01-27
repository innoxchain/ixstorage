package event

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"fmt"
	"time"
	"github.com/satori/go.uuid"
)


func TestDeserializeOrderRevisedEvent(t *testing.T) {
	orderRevisedEvent := &OrderRevisedEvent{
		RevisedBy: "me",
		Reason:    "not enough money",
	}

	registerEvent(orderRevisedEvent)

	pe := PersistentEvent{
		AggregateID:   uuid.NewV4().String(),
		AggregateType: "order",
		EventType:     "OrderRevisedEvent",
		CreatedAt:     time.Now().String(),
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

	orderRevisedEvent := OrderRevisedEvent{
		RevisedBy: "me",
		Reason:    "not enough money",
	}

	event := Event{
		AggregateID:   uuid.NewV4().String(),
		AggregateType: "order",
		EventType:     "OrderRevisedEvent",
		CreatedAt:     time.Now().String(),
		Payload:	   orderRevisedEvent,
	}

	persistentEvent, err := event.Serialize()

	if err != nil {
		fmt.Println("couldn't serialize Event: ", event)
	}

	expected := "{\"RevisedBy\":\"me\",\"Reason\":\"not enough money\"}"

	assert.Equal(t, expected, persistentEvent.RawData, "Expected result doesn't match!")
}