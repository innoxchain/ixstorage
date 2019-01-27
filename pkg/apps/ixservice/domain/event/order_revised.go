package event

import (
	"time"
)

//OrderRevisedEvent is the event initiated when an existing order has been revised
type OrderRevisedEvent struct {
	RevisedBy string
	Reason    string
}

func (OrderRevisedEvent) GetAggregateType() string {
	return "order"
}

func (OrderRevisedEvent) GetEventType() string {
	return "OrderRevisedEvent"
}

func (OrderRevisedEvent) GetCreatedAt() string {
	return time.Now().String()
}
