package event

import (
	"fmt"
	"time"
)

//////////////////////////////////////////////////////////////////////
// Order Aggregate Definition 										//
//////////////////////////////////////////////////////////////////////

type Order struct {
	BaseAggregate
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
	confirmedBy: %s
	revisedBy: %s
	revisedReason: %s
	(Pending Changes: %d)
	(Version: %d)`

	return fmt.Sprintf(format, o.UUID, o.LastModified, o.Capacity, o.ConfirmedBy, o.RevisedStat.RevisedBy, o.RevisedStat.Reason, len(o.Changes), o.Version)
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

func (oc OrderCreated) Apply(aggregate Aggregate, event *Event) {
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

func (oc OrderConfirmed) Apply(aggregate Aggregate, event *Event) {
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

func (or OrderRevised) Apply(aggregate Aggregate, event *Event) {
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

func (coc CreateOrderCommand) CreateBaseEvent() BaseEvent {
	return OrderCreated{
		UUID:     coc.UUID,
		Capacity: coc.Capacity,
	}
}

type ConfirmOrderCommand struct {
	ConfirmedBy string
}

func (coc ConfirmOrderCommand) CreateBaseEvent() BaseEvent {
	return OrderConfirmed{
		ConfirmedBy: coc.ConfirmedBy,
	}
}

type ReviseOrderCommand struct {
	RevisedBy string
	Reason string
}

func (roc ReviseOrderCommand) CreateBaseEvent() BaseEvent {
	return OrderRevised{
		RevisedBy: roc.RevisedBy,
		Reason: roc.Reason,
	}
}

//////////////////////////////////////////////////////////////////////
// Order Enums														//
//////////////////////////////////////////////////////////////////////

//Capacity is the capacity in GB available for customers
type Capacity int

const (
	//ThreeGB allows 3 GB of capacity
	ThreeGB Capacity = iota
	//SixGB allows 6 GB of capacity
	SixGB Capacity = iota
	//TenGB allows 10 GB of capacity
	TenGB Capacity = iota
)

func (c Capacity) String() string {
	return [...]string{"ThreeGB", "SixGB", "TenGB"}[c]
}