package event

//OrderConfirmedEvent is the event initiated when an existing order has been confirmed
type OrderConfirmedEvent struct {
	Event
	ConfirmedBy string
}
