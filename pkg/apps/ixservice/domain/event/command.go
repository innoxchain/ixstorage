package event

import (
	log "github.com/sirupsen/logrus"
	//store "github.com/innoxchain/ixstorage/pkg/apps/ixservice/eventstore"
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
}