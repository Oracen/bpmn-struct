package foundation

import "github.com/Oracen/bpmn-struct/shared"

// Abstract super class for all elements
type RootElement struct {
	BaseElement
}

func CreateRootElement(id string) RootElement {
	baseElement := CreateBaseElement(id)
	return RootElement{
		BaseElement: baseElement,
	}
}

func (r RootElement) Validate(name string) []error {
	name = shared.TypeNameString(name, r, r.Id)
	return r.BaseElement.Validate(name)
}
