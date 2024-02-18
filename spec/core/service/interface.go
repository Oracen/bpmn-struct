package service

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/spec/process/activity"
	"github.com/Oracen/bpmn-struct/validation"
)

type Interface struct {
	foundation.RootElement
	Name             string                     `xml:"name" json:"name"`
	Operations       []Operation                `xml:"operations" json:"operations"`
	CallableElements []activity.CallableElement `xml:"callableElements" json:"callableElements"`
	Element          []any                      `xml:"element" json:"element"`
}

func CreateInterface(id, name string, operations []Operation) Interface {
	rootElement := foundation.CreateRootElement(id)
	return Interface{
		RootElement:      rootElement,
		Name:             name,
		Operations:       operations,
		CallableElements: []activity.CallableElement{},
		Element:          []any{},
	}
}

func (i Interface) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, i, i.Id)
	checks = append(checks, i.RootElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, i.Operations)...)
	checks = append(checks, validation.ArrCheckItems(name, i.CallableElements)...)
	checks = append(
		checks,
		validation.ArrOneOrMore(name, "Operations", i.Operations),
		validation.ArrZeroOne(name, "Element", i.Element),
		validation.ValNonzero(name, "Name", i.Name),
	)
	return validation.FilterErrors(checks)
}
