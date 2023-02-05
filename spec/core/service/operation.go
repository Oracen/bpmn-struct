package service

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/common"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type Operation struct {
	foundation.BaseElement
	Name              string           `xml:"name" json:"name"`
	InMessageRef      common.Message   `xml:"inMessageRef" json:"inMessageRef"`
	OutMessageRef     []common.Message `xml:"outMessageRef" json:"outMessageRef"`
	ErrorRef          []common.Error   `xml:"errorRef" json:"errorRef"`
	ImplementationRef []any            `xml:"implementationRef" json:"implementationRef"`
}

func CreateOperation(id, name string, inMessage common.Message) Operation {
	baseElement := foundation.CreateBaseElement(id)
	return Operation{
		BaseElement:       baseElement,
		Name:              name,
		InMessageRef:      inMessage,
		OutMessageRef:     []common.Message{},
		ErrorRef:          []common.Error{},
		ImplementationRef: []any{},
	}
}

func (o Operation) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, o, o.Id)
	checks = append(checks, o.BaseElement.Validate(name)...)
	checks = append(checks, o.InMessageRef.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, o.OutMessageRef)...)
	checks = append(checks, validation.ArrCheckItems(name, o.ErrorRef)...)
	checks = append(
		checks,
		validation.ArrZeroOne(name, "OutMessageRef", o.OutMessageRef),
		validation.ArrZeroOne(name, "ImplementationRef", o.ImplementationRef),
		validation.ValNonzero(name, "Name", o.Name),
	)

	return validation.FilterErrors(checks)
}
