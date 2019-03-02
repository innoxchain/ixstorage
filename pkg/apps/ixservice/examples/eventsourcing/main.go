package main

import (
	"github.com/satori/go.uuid"
	//"fmt"
	//"time"
	es "github.com/innoxchain/ixstorage/pkg/apps/ixservice/event"
)

/*
//////////////////////////////////////////////////////////////////////
// Order Aggregate Definition 										//
//////////////////////////////////////////////////////////////////////

type Order struct {
	event.BaseAggregate
	Capacity    Capacity
	ConfirmedBy string
	RevisedStat RevisedStatus
}

type RevisedStatus struct {
	RevisedBy string
	Reason    string
}

func (o *Order) String() string {
	format := `Order:
	uuid: %s
	lastModified: %s
	capacity: %s
	(Pending Changes: %d)
	(Version: %d)`

	return fmt.Sprintf(format, o.UUID, o.LastModified, o.Capacity, len(o.Changes), o.Version)
}


//////////////////////////////////////////////////////////////////////
// Order Event Definitions 											//
//////////////////////////////////////////////////////////////////////

type OrderCreated struct {
	UUID     string		`json:"uuid"`
	Capacity Capacity  	`json:"capacity"`
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
	return time.Now().UTC()
}

func (oc OrderCreated) Apply(aggregate event.Aggregate, event *event.Event) {
	order := aggregate.(*Order)
	order.UUID = oc.UUID
	order.Capacity = oc.Capacity
	event.AggregateID = oc.UUID
}

type OrderConfirmed struct {
	ConfirmedBy string `json:"confirmedBy"`
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
	return time.Now().UTC()
}

func (oc OrderConfirmed) Apply(aggregate event.Aggregate, event *event.Event) {
	order := aggregate.(*Order)
	order.ConfirmedBy = oc.ConfirmedBy
}

type OrderRevised struct {
	RevisedBy string `json:"revisedBy"`
	Reason    string `json:"reason"`
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
	return time.Now().UTC()
}

func (or OrderRevised) Apply(aggregate event.Aggregate, event *event.Event) {
	order := aggregate.(*Order)
	order.RevisedStat = RevisedStatus{RevisedBy: or.RevisedBy, Reason: or.Reason}
}


//////////////////////////////////////////////////////////////////////
// Order Command Definition											//
//////////////////////////////////////////////////////////////////////

type CreateOrderCommand struct {
	UUID string
	Capacity Capacity
}

func (coc CreateOrderCommand) CreateBaseEvent() event.BaseEvent {
	return OrderCreated{
		UUID:     coc.UUID,
		Capacity: coc.Capacity,
	}
}
*/

func main() {

	var order es.Order

	es.RegisterEvent(&es.OrderCreated{})
	es.RegisterEvent(&es.OrderConfirmed{})
	es.RegisterEvent(&es.OrderRevised{})


	createCommand := es.CreateOrderCommand{
		UUID: uuid.NewV4().String(),
		Capacity: es.ThreeGB,
	}

	confirmCommand := es.ConfirmOrderCommand{
		ConfirmedBy: "myself",
	}

	es.ApplyCommand(createCommand, &order)
	es.ApplyCommand(confirmCommand, &order)

}