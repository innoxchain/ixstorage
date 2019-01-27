package event

import (
	//"fmt"
	//"reflect"
	"testing"
	//"time"

	//"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	//"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
	//log "github.com/sirupsen/logrus"
)

func TestDomainEventRegisterEvents(t *testing.T) {
	//aggregateID := uuid.NewV4().String()
	//capacity := enum.SixGB
	//issuer := "me"

	//orderCreatedEvent := &OrderCreatedEvent{Event: Event{AggregateID: aggregateID, EventType: "OrderCreated" ,CreatedAt: time.Now().String()}, Capacity: capacity}
	orderCreatedEvent := &OrderCreatedEvent{}
	//orderConfirmedEvent := &OrderConfirmedEvent{Event: Event{AggregateID: aggregateID, EventType: "OrderConfirmed", CreatedAt: time.Now().String()}, ConfirmedBy: issuer}
	orderConfirmedEvent := &OrderConfirmedEvent{}

	registerEvent(orderCreatedEvent)
	registerEvent(orderConfirmedEvent)

	assert.Equal(t, "OrderCreatedEvent", getEvent("OrderCreatedEvent").Name(), "Registered Event should be event.OrderCreatedEvent")
	assert.Equal(t, "OrderConfirmedEvent", getEvent("OrderConfirmedEvent").Name(), "Registered Event should be event.OrderConfirmedEvent")

	/*
		fmt.Println(getEvent("OrderCreatedEvent"))
		fmt.Println(getEvent("OrderConfirmedEvent"))

		fmt.Printf("%T\n", makeInstance("OrderCreatedEvent"))
		fmt.Printf("%T\n", makeInstance("OrderConfirmedEvent"))
	*/
}

/*
func TestOrderRevisedEvent(t *testing.T) {
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
		fmt.Println("couldn't deserialize DataStoreEvent: ", pe)
	}

	v := reflect.ValueOf(event.Payload)
	kv := v.MapKeys()
	strct := v.MapIndex(kv[0])

	assert.Equal(t, strct.Interface().(string), "me", "RevisedBy must be \"me\"")

	strct = v.MapIndex(kv[1])
	assert.Equal(t, strct.Interface().(string), "not enough money", "Reason must be \"not enough money\"")


	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			log.Info("key: ", key.Interface(), " value: ", strct.Interface())
			//fmt.Println(key.Interface(), strct.Interface())
		}
	}

	//assert.Nil(t, event.Payload, "Event Payload must not be nil")
}
*/
