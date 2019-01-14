package event

import(
	"encoding/json"
)

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

func (e *OrderConfirmedEvent) MarshalJSON() (b []byte, err error) {  
    return json.Marshal(map[string]string{
		"AggregateId":  e.GetAggregateID(),
		"EventType": e.GetType(),
		"CreatedAt": e.GetCreatedAt(),
		"ConfirmedBy": e.ConfirmedBy,
    })
}
