package event

import (
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
