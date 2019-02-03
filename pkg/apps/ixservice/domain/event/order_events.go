package event

import (
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
)

/*
import (
	"time"
)
*/

//OrderRevised is the event payload initiated when an existing order has been revised
type OrderRevised struct {
	RevisedBy string
	Reason    string
}

func (OrderRevised) GetAggregateType() string {
	return "order"
}

func (OrderRevised) GetEventType() string {
	return "OrderRevised"
}

func (OrderRevised) GetSequence() int {
	return 3
}

/*
func (OrderRevised) GetCreatedAt() string {
	return time.Now().String()
}
*/


type OrderCreated struct {
	Capacity enum.Capacity
}

func (OrderCreated) GetAggregateType() string {
	return "order"
}

func (OrderCreated) GetEventType() string {
	return "OrderCreated"
}

func (OrderCreated) GetSequence() int {
	return 1
}


type OrderConfirmed struct {
	ConfirmedBy string
}

func (OrderConfirmed) GetAggregateType() string {
	return "order"
}

func (OrderConfirmed) GetEventType() string {
	return "OrderConfirmed"
}

func (OrderConfirmed) GetSequence() int {
	return 2
}