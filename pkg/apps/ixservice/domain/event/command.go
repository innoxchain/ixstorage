package event

import (
	log "github.com/sirupsen/logrus"
)

type Command interface {
	CreateBaseEvent() BaseEvent
}

func ApplyCommand(command Command, aggregate Aggregate) {
	
	baseEvent:=command.CreateBaseEvent()

	RegisterEvent(baseEvent)

	event:=BuildEvent(baseEvent, aggregate.GetAggregateID())

	event.ApplyChanges(aggregate)

	log.Info("AGGREGATE AFTER COMMAND WAS APPLIED: ", aggregate)
	
	/*
	event.RegisterEvent(orderConfirmed)
	e = event.BuildEvent(orderConfirmed, order.UUID)
	e.ApplyChanges(&order)
	db.Persist(&order.BaseAggregate)
	*/
}