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

	assert.Equal(t, orderCreatedEvent, getEvent("OrderCreatedEvent"), "Registered Event should be event.OrderCreatedEvent")
	assert.Equal(t, orderConfirmedEvent, getEvent("OrderConfirmedEvent"), "Registered Event should be event.OrderConfirmedEvent")
}
