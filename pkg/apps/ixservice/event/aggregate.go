package event

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type Aggregate interface {
	GetAggregateID() string
	GetVersion() int
	trackChanges(e Event)
	incrementVersion()
	GetChanges() []Event
	MarkAsCommited()
}

type BaseAggregate struct {
	UUID         string    `json:"uuid"`
	Version      int       `json:"version"`
	LastModified time.Time `json:"lastModified"`
	Changes      []Event   `json:"changes"`
}

func (a *BaseAggregate) GetChanges() []Event {
	return a.Changes
}

func (a *BaseAggregate) GetAggregateID() string {
	return a.UUID
}

func (a *BaseAggregate) GetVersion() int {
	return a.Version
}

func (a *BaseAggregate) incrementVersion() {
	a.Version += 1
}

func (a *BaseAggregate) trackChanges(e Event) {
	a.LastModified = e.CreatedAt
	a.Changes = append(a.Changes, e)
}

func (a *BaseAggregate) MarkAsCommited() {
	a.Changes = nil
	a.Changes = make([]Event, 0)
}

func Replay(aggregate Aggregate, events []Event) {
	for _, ev := range events {
		ev.ApplyChanges(aggregate)
	}
}

func (a *BaseAggregate) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage

	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	var rawEventMessages []*json.RawMessage
	err = json.Unmarshal(*objMap["changes"], &rawEventMessages)
	if err != nil {
		return err
	}

	convertRawToAggregate(a, objMap)

	a.Changes = make([]Event, len(rawEventMessages))

	var m map[string]string

	for i, rawMessage := range rawEventMessages {
		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}

		e := Event{}
		convertStringToEvent(&e, m)

		a.Changes[i] = e
	}

	return nil
}

func convertStringToEvent(event *Event, m map[string]string) {
	createdAt := m["CreatedAt"]
	modifiedDate, _ := time.Parse(time.RFC3339, createdAt)
	aggregateID := m["AggregateID"]
	aggregateType := m["AggregateType"]
	eventType := m["EventType"]
	payload := m["Payload"]
	sequence, _ := strconv.Atoi(string(m["Sequence"]))

	event.AggregateID = aggregateID
	event.AggregateType = aggregateType
	event.CreatedAt = modifiedDate
	event.EventType = eventType
	event.Sequence = sequence

	/*
	dataPointer := reflect.New(eventRegistry[eventType])
	dataValue := dataPointer.Elem()
	iface := dataValue.Interface()
	*/

	objType := reflect.TypeOf(getEvent(eventType)).Elem()
    obj := reflect.New(objType).Interface()

	//err := json.Unmarshal([]byte(payload), &iface)
	err := json.Unmarshal([]byte(payload), &obj)
	if err != nil {
		log.Fatal(err)
	}

	//event.Payload = iface
	event.Payload = obj
}

func convertRawToAggregate(aggregate *BaseAggregate, objMap map[string]*json.RawMessage) {
	timeString, _ := strconv.Unquote(string(*objMap["lastModified"]))
	aggregate.LastModified, _ = time.Parse(time.RFC3339, timeString)
	aggregate.UUID, _ = strconv.Unquote(string(*objMap["uuid"]))
	aggregate.Version, _ = strconv.Atoi(string(*objMap["version"]))
}
