package foundation

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/validation"
)

type BaseElement struct {
	Id                   string                    `xml:"id" json:"id"`
	Documentation        []Documentation           `xml:"documentation" json:"documentation"`
	ExtensionDefinitions []ExtensionDefinition     `xml:"extensionDefinition" json:"extensionDefinition"`
	ExtensionValues      []ExtensionAttributeValue `xml:"extensionValues" json:"extensionValues"`
}

func CreateBaseElement(id string) BaseElement {
	return BaseElement{
		Id:                   id,
		Documentation:        []Documentation{},
		ExtensionDefinitions: []ExtensionDefinition{},
		ExtensionValues:      []ExtensionAttributeValue{},
	}
}

func (b BaseElement) Validate(name string) (errors []error) {
	checks := []error{}
	name = shared.TypeNameString(name, b, b.Id)
	checks = append(checks, validation.ArrCheckItems(name, b.Documentation)...)
	checks = append(checks, validation.ArrCheckItems(name, b.ExtensionDefinitions)...)
	checks = append(checks, validation.ArrCheckItems(name, b.ExtensionValues)...)
	checks = append(checks, validation.ValNonzero(name, "Id", b.Id))
	return validation.FilterErrors(checks)
}
