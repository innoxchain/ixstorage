package event

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
	"github.com/stretchr/testify/assert"
)

func TestOrderCreatedCommand(t *testing.T) {

	var order Order

	RegisterEvent(&OrderCreated{})

	createCommand := CreateOrderCommand{
		UUID: uuid.NewV4().String(),
		Capacity: enum.ThreeGB,
	}

	ApplyCommand(createCommand, &order)

	t.Log("Order UUID: ", order.GetAggregateID())

	assert.True(t, order.UUID!="", "AggregateID must not be nil")
	assert.Equal(t, enum.ThreeGB, order.Capacity, "Capacity must be ThreeGB after command has been applied")
}
