package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainEventRegisterEvents(t *testing.T) {
	orderCreatedEvent := &OrderCreatedEvent{}
	orderConfirmedEvent := &OrderConfirmedEvent{}

	RegisterEvent(orderCreatedEvent)
	RegisterEvent(orderConfirmedEvent)

	assert.Equal(t, "OrderCreatedEvent", getEvent("OrderCreatedEvent").Name(), "Registered Event should be event.OrderCreatedEvent")
	assert.Equal(t, "OrderConfirmedEvent", getEvent("OrderConfirmedEvent").Name(), "Registered Event should be event.OrderConfirmedEvent")
}