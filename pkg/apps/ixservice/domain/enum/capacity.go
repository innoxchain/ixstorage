package enum


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