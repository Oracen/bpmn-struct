package foundation

import (
	"bpmn-struct/validation"
	"fmt"
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
	if name == "" {
		name = fmt.Sprintf("Import:%s", b.Id)
	}

	checks := []error{
		validation.ValNonzero(name, "Id", b.Id),
	}
	return validation.FilterErrors(checks)
}
