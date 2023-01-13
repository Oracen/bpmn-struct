package foundation

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
	return r.BaseElement.Validate(name)
}
