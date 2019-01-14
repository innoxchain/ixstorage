package enum

//OrderStatus is showing the possible event states of an aggregate
type OrderStatus int

const (
	//Created is set when a new order has been created
	Created OrderStatus = iota
	//Confirmed is set when an existing order has been confirmed
	Confirmed OrderStatus = iota
)

func (os OrderStatus) String() string {
	return [...]string{"Created", "Confirmed"}[os]
}