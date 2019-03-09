package event

import (
	log "github.com/sirupsen/logrus"
)

type Command interface {
	CreateBaseEvent() BaseEvent
}

func ApplyCommand(command Command, aggregate Aggregate) {

	db := EventStore{}
	
	baseEvent:=command.CreateBaseEvent()

	event:=BuildEvent(baseEvent, aggregate.GetAggregateID())

	event.ApplyChanges(aggregate)

	log.Info("AGGREGATE AFTER COMMAND WAS APPLIED: ", aggregate)

	db.Persist(aggregate)

	log.Info("AGGREGATE AFTER IT WAS PERSISTED TO DATABASE: ", aggregate)
}