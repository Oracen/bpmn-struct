package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type Escalation struct {
	foundation.RootElement
	StructureRef   []ItemDefinition `xml:"structureRef" json:"structureRef"`
	Name           string           `xml:"name" json:"name"`
	EscalationCode string           `xml:"escalationCode" json:"escalationCode"`
}

func CreateEscalation(id, name, escalationCode string) Escalation {
	rootElement := foundation.CreateRootElement(id)
	return Escalation{
		RootElement:    rootElement,
		StructureRef:   []ItemDefinition{},
		Name:           name,
		EscalationCode: escalationCode,
	}
}

func (e Escalation) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, e, e.Id)
	checks = append(checks, e.RootElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, e.StructureRef)...)
	checks = append(checks, validation.ArrZeroOne(name, "StructureRef", e.StructureRef))
	checks = append(
		checks,
		validation.ValNonzero(name, "Name", e.Name),
		validation.ValNonzero(name, "EscalationCode", e.EscalationCode),
	)
	return validation.FilterErrors(checks)
}
