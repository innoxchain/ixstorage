package event

//OrderConfirmedEvent is the event initiated when an existing order has been confirmed
type OrderConfirmedEvent struct {
	AggregateID string
	CreatedAt   string
	ConfirmedBy string
}

func (e *OrderConfirmedEvent) GetType() string {
	return "order.confirmed"
}

func (e *OrderConfirmedEvent) GetAggregateID() string {
	return e.AggregateID
}

func (e *OrderConfirmedEvent) GetCreatedAt() string {
	return e.CreatedAt
}
