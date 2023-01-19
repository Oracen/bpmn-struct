package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/validation"
)

type Event struct {
	FlowNode
}

func CreateEvent(id string) Event {
	flowNode := CreateFlowNode(id)
	return Event{
		FlowNode: flowNode,
	}
}

func (e Event) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, e, e.Id)
	checks = append(checks, e.FlowElement.Validate(name)...)
	return validation.FilterErrors(checks)
}
