package event

import (
	"reflect"
	"encoding/json"
)

var eventRegistry = make(map[string]reflect.Type)

type DomainEvent interface {
	GetEventType() string
	GetAggregateID() string
	GetCreatedAt() string
}

type Event struct {
	AggregateID string
	EventType	string
	CreatedAt   string
}

func (e *Event) GetAggregateID() string {
	return e.AggregateID
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
		"AggregateId":  e.GetAggregateID(),
		"EventType": e.GetEventType(),
		"CreatedAt": e.GetCreatedAt(),
    })
}