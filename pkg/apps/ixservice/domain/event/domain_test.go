package event

import (
	//"fmt"
	//"time"
	"testing"

	//"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	//"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
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