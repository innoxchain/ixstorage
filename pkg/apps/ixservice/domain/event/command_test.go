package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
)

func TestOrderCreatedCommand(t *testing.T) {
	var order Order

	createCommand:=CreateOrderCommand{
		Capacity: enum.ThreeGB,
	}

	ApplyCommand(createCommand,&order)

	assert.Equal(t, enum.ThreeGB, order.Capacity, "Capacity must be ThreeGB after command has been applied")
}