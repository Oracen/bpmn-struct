package common

type FlowElement struct {
}

func CreateFlowElement() FlowElement {
	return FlowElement{}
}

func (f FlowElement) Validate(name string) []error {
	return []error{}
}
