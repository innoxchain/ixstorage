package event

type Event struct {
	aggregateID string
	eventType	string
	createdAt   string
}

func (e *Event) GetAggregateID() string {
	return e.aggregateID
}

func (e *Event) GetEventType() string {
	return e.eventType
}

func (e *Event) GetCreatedAt() string {
	return e.createdAt
}

func (e *Event) SetEventType(eventType string) {
	e.eventType = eventType
}

func (e *Event) SetAggregateID(aggregateID string) {
	e.aggregateID = aggregateID
}

func (e *Event) SetCreatedAt(createdAt string) {
	e.createdAt = createdAt
}