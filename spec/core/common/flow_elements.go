package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/spec/process"
	"github.com/Oracen/bpmn-struct/validation"
)

type FlowElement struct {
	foundation.BaseElement
	Name             []string             `xml:"name" json:"name"`
	CategoryValueRef []CategoryValue      `xml:"categoryValueRef" json:"categoryValueRef"`
	Auditing         []process.Auditing   `xml:"auditing" json:"auditing"`
	Monitoring       []process.Monitoring `xml:"monitoring" json:"monitoring"`
}

func CreateFlowElement(id string) FlowElement {
	baseElement := foundation.CreateBaseElement(id)
	return FlowElement{
		BaseElement: baseElement,
	}
}

func (f FlowElement) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, f, f.Id)
	checks = append(checks, f.BaseElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, f.Auditing)...)
	checks = append(checks, validation.ArrCheckItems(name, f.Monitoring)...)
	checks = append(
		checks,
		validation.ArrZeroOne(name, "Name", f.Name),
		validation.ArrZeroOne(name, "Auditing", f.Auditing),
		validation.ArrZeroOne(name, "Monitoring", f.Monitoring),
	)
	return validation.FilterErrors(checks)
}

type FlowNode struct {
	FlowElement
}

func CreateFlowNode(id string) FlowNode {
	flowElement := CreateFlowElement(id)
	return FlowNode{
		FlowElement: flowElement,
	}
}

func (f FlowNode) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, f, f.Id)
	checks = append(checks, f.FlowElement.Validate(name)...)
	return validation.FilterErrors(checks)
}
