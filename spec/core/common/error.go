package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type Error struct {
	foundation.RootElement
	StructureRef []ItemDefinition `xml:"structureRef" json:"structureRef"`
	Name         string           `xml:"name" json:"name"`
	ErrorCode    string           `xml:"errorCode" json:"errorCode"`
}

func CreateError(id, name, errorCode string) Error {
	rootElement := foundation.CreateRootElement(id)
	return Error{
		RootElement:  rootElement,
		StructureRef: []ItemDefinition{},
		Name:         name,
		ErrorCode:    errorCode,
	}
}

func (e Error) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, e, e.Id)
	checks = append(checks, e.RootElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, e.StructureRef)...)
	checks = append(checks, validation.ArrZeroOne(name, "StructureRef", e.StructureRef))
	return validation.FilterErrors(checks)
}
