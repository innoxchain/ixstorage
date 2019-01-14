package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/enum"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

//Order is the Aggregate which will control the behaviour of state transitions of this Aggregate.
type Order struct {
	UUID    string              `json:"uuid"`
	Version int                 `json:"version"`
	Status  enum.OrderStatus    `json:"status"`
	Changes []event.DomainEvent `json:"changes"`
}

func (o *Order) createOrder(cap enum.Capacity) event.DomainEvent {
	//Todo: store new event in database from a new service
	o.Version = 0
	event := &event.OrderCreatedEvent{AggregateID: uuid.NewV4().String(), CreatedAt: time.Now().String(), Capacity: cap}
	o.trackChange(event)

	return event
}

func (o *Order) confirmOrder(issuer string) event.DomainEvent {
	//Todo: store new event in database from a new service
	event := &event.OrderConfirmedEvent{AggregateID: o.UUID, CreatedAt: time.Now().String(), ConfirmedBy: issuer}
	o.trackChange(event)

	return event
}

func (o *Order) trackChange(event event.DomainEvent) {
	o.Changes = append(o.Changes, event)
	o.transition(event)
}

func (o *Order) transition(ev event.DomainEvent) {
	switch e := ev.(type) {
	case *event.OrderCreatedEvent:
		o.UUID = e.GetAggregateID()
		o.Status = enum.Created
	case *event.OrderConfirmedEvent:
		o.UUID = e.GetAggregateID()
		o.Status = enum.Confirmed
	}
	o.Version++
}

func GetOrderFromHistory(events []event.DomainEvent, o *Order) {
	//order := &Order{}
	for _, ev := range events {
		o.transition(ev)
	}
	//return o
}

func MarkOrderAsCommitted(o *Order) {
	o.Changes = nil
	o.Changes = make([]event.DomainEvent, 0)
}

func (o *Order) String() string {
	format := `Order:
	uuid: %s
	status: %s
	(Pending Changes: %d)
	(Version: %d)`

	return fmt.Sprintf(format, o.UUID, o.Status, len(o.Changes), o.Version)
}

func (o *Order) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage

	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	log.Info("objMap: ", objMap)

	var rawOrderMessages []*json.RawMessage
	err = json.Unmarshal(*objMap["changes"], &rawOrderMessages)
	if err != nil {
		return err
	}

	log.Info("raw message: ", rawOrderMessages)

	o.UUID, _ = strconv.Unquote(string(*objMap["uuid"]))
	o.Version, _ = strconv.Atoi(string(*objMap["version"]))
	o.Changes = make([]event.DomainEvent, len(rawOrderMessages))

	var m map[string]string
	for i, rawMessage := range rawOrderMessages {
		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}
		log.Info("deserialized raw message: ", m)

		if m["EventType"] == "order.created" {
			var e event.OrderCreatedEvent
			err := json.Unmarshal(*rawMessage, &e)
			if err != nil {
				log.Fatal(err)
			}
			o.Changes[i] = &e
		} else if m["EventType"] == "order.confirmed" {
			var e event.OrderConfirmedEvent
			err := json.Unmarshal(*rawMessage, &e)
			if err != nil {
				log.Fatal(err)
			}
			o.Changes[i] = &e
		} else {
			return errors.New("Unsupported EventType")
		}
	}

	return nil
}
