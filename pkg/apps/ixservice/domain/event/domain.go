package event

import (
	"encoding/json"
	"reflect"
)

var eventRegistry = make(map[string]reflect.Type)

type DomainEvent interface {
	GetEventType() string
	GetAggregateType() string
	GetCreatedAt() string
}

type Event struct {
	AggregateID   string
	AggregateType string
	EventType     string
	CreatedAt     string
	Payload       interface{}
}

type PersistentEvent struct {
	AggregateID   string
	AggregateType string
	EventType     string
	CreatedAt     string
	RawData       string
}

func (e *Event) GetAggregateType() string {
	return e.AggregateType
}

func (e *Event) GetEventType() string {
	return e.EventType
}

func (e *Event) GetCreatedAt() string {
	return e.CreatedAt
}

func registerEvent(event interface{}) {
	t := reflect.TypeOf(event).Elem()
	eventRegistry[t.Name()] = t
}

func getEvent(key string) reflect.Type {
	return eventRegistry[key]
}

func makeInstance(name string) interface{} {
	return reflect.New(eventRegistry[name]).Elem().Interface()
}

func (e *Event) MarshalJSON() (b []byte, err error) {

	return json.Marshal(map[string]string{
		"AggregateType": e.GetAggregateType(),
		"EventType":     e.GetEventType(),
		"CreatedAt":     e.GetCreatedAt(),
	})
}

func (e Event) Serialize() (PersistentEvent, error) {
	var err error
	result := PersistentEvent{}

	result.AggregateID = e.AggregateID
	result.AggregateType = e.AggregateType
	result.CreatedAt = e.CreatedAt
	result.EventType = e.EventType
	
	ser, err := json.Marshal(e.Payload)
	if err != nil {
		return PersistentEvent{}, err
	}

	result.RawData = string(ser)

	return result, nil
}

func (e PersistentEvent) Deserialize() (Event, error) {
	var err error
	result := Event{}

	eventType := e.EventType
	dataPointer := reflect.New(eventRegistry[eventType])
	dataValue := dataPointer.Elem()
	iface := dataValue.Interface()

	err = json.Unmarshal([]byte(e.RawData), &iface)
	if err != nil {
		return Event{}, err
	}

	result.AggregateID = e.AggregateID
	result.AggregateType = e.AggregateType
	result.EventType = e.EventType
	result.CreatedAt = e.CreatedAt
	result.Payload = iface

	return result, nil
}
