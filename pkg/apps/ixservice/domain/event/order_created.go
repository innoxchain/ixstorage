package event

import(
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
)

//OrderCreatedEvent is the event initiated when a new order has been created
type OrderCreatedEvent struct {
	Event
	Capacity enum.Capacity
}