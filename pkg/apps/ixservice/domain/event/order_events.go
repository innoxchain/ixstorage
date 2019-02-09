package event

import (
	"time"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
)

type Order struct {
	BaseAggregate
}

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

func (OrderCreated) GetCreatedAt() time.Time {
	return time.Now()
}

func (OrderCreated) Apply(aggregate Aggregate, event Event) {
	//order := aggregate.(*Order)
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

func (OrderConfirmed) GetCreatedAt() time.Time {
	return time.Now()
}

func (OrderConfirmed) Apply(aggregate Aggregate, event Event) {
	//order := aggregate.(*Order)
}


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

func (OrderRevised) GetCreatedAt() time.Time {
	return time.Now()
}

func (OrderRevised) Apply(aggregate Aggregate, event Event) {
	//order := aggregate.(*Order)
}