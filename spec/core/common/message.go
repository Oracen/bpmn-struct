package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type Message struct {
	foundation.RootElement
	ItemRef []ItemDefinition `xml:"itemRef" json:"itemRef"`
}

func CreateMessage(id string) Message {
	rootElement := foundation.CreateRootElement(id)
	return Message{
		RootElement: rootElement,
		ItemRef:     []ItemDefinition{},
	}
}

func (m Message) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, m, m.Id)
	checks = append(checks, m.RootElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, m.ItemRef)...)
	checks = append(
		checks,
		validation.ArrZeroOne(name, "ItemRef", m.ItemRef),
	)

	return validation.FilterErrors(checks)
}
