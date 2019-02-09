package event

import (
	//"strconv"
	"time"
	"encoding/json"
	"reflect"
	log "github.com/sirupsen/logrus"
)

var eventRegistry = make(map[string]reflect.Type)

type DomainEvent interface {
	GetEventType() string
	GetAggregateType() string
	GetCreatedAt() time.Time
	GetSequence() int
	//Apply(aggregate Aggregate, event Event)
}

type BaseEvent interface {
	GetEventType() string
	GetAggregateType() string
	GetCreatedAt() time.Time
	GetSequence() int
	Apply(aggregate Aggregate, event Event)
}

type Event struct {
	AggregateID   string
	AggregateType string
	EventType     string
	CreatedAt     time.Time
	Sequence	  int
	Payload       interface{}
}

type PersistentEvent struct {
	AggregateID   string
	AggregateType string
	EventType     string
	CreatedAt     time.Time
	Sequence	  int
	RawData       string
}

func (e *Event) GetAggregateType() string {
	return e.AggregateType
}

func (e *Event) GetEventType() string {
	return e.EventType
}

func (e *Event) GetSequence() int {
	return e.Sequence
}

func (e *Event) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e Event) apply(agg Aggregate) {
	e.Payload.(BaseEvent).Apply(agg, e)
	agg.incrementVersion()
	agg.trackChanges(e)
}

func RegisterEvent(event interface{}) {
	log.Info("RegisterEvent: ", event)

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
		//"Sequence": 	 strconv.Itoa(e.GetSequence()),
		//"CreatedAt":     e.GetCreatedAt().String(),
	})
}

func BuildEvent(de DomainEvent, aggregateID string) Event {
	event := Event{}

	event.AggregateID=aggregateID
	event.AggregateType=de.GetAggregateType()
	event.EventType=de.GetEventType()
	event.Sequence=de.GetSequence()
	event.CreatedAt=de.GetCreatedAt()
	event.Payload=de

	return event
}

func (e Event) Serialize() (PersistentEvent, error) {
	var err error
	result := PersistentEvent{}

	result.AggregateID = e.AggregateID
	result.AggregateType = e.AggregateType
	result.CreatedAt = e.CreatedAt
	result.EventType = e.EventType
	result.Sequence = e.Sequence
	
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
	result.Sequence = e.Sequence
	result.Payload = iface

	return result, nil
}
