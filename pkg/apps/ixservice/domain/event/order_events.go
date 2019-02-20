package event

import (
	"github.com/satori/go.uuid"
	//"strconv"
	//"encoding/json"
	"fmt"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
	"time"
)

type RevisedStatus struct {
	RevisedBy string
	Reason    string
}

type Order struct {
	BaseAggregate
	Capacity    enum.Capacity
	ConfirmedBy string
	RevisedStat RevisedStatus
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

type CreateOrderCommand struct {
	Capacity enum.Capacity
}

func (coc CreateOrderCommand) CreateBaseEvent() BaseEvent {
	return OrderCreated{
		UUID:     uuid.NewV4().String(),
		Capacity: coc.Capacity,
	}
}

type OrderCreated struct {
	UUID     string			`json:"uuid"`
	Capacity enum.Capacity  `json:"capacity"`
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

func (oc OrderCreated) Apply(aggregate Aggregate, event Event) {
	order := aggregate.(*Order)
	order.UUID = oc.UUID
	order.Capacity = oc.Capacity
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

func (oc OrderConfirmed) Apply(aggregate Aggregate, event Event) {
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

func (or OrderRevised) Apply(aggregate Aggregate, event Event) {
	order := aggregate.(*Order)
	order.RevisedStat = RevisedStatus{RevisedBy: or.RevisedBy, Reason: or.Reason}
}