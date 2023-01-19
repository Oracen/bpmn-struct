package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/spec/process"
	"github.com/Oracen/bpmn-struct/validation"
)

type FlowElementsContainer struct {
	foundation.BaseElement
	FlowElements []FlowElement     `xml:"flowElements" json:"flowElements"`
	LaneSets     []process.LaneSet `xml:"laneSets" json:"laneSets"`
}

func CreateFlowElementsContainer(id string) FlowElementsContainer {
	baseElement := foundation.CreateBaseElement(id)
	return FlowElementsContainer{
		BaseElement:  baseElement,
		FlowElements: []FlowElement{},
		LaneSets:     []process.LaneSet{},
	}
}

func (f FlowElementsContainer) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, f, f.Id)
	checks = append(checks, f.BaseElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, f.FlowElements)...)
	checks = append(checks, validation.ArrCheckItems(name, f.LaneSets)...)
	return validation.FilterErrors(checks)
}
