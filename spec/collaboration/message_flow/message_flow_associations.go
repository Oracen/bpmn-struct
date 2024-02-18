package message_flow

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type MessageFlowAssociation struct {
	foundation.BaseElement
	InnerMessageFlowRef MessageFlow `xml:"innerMessageFlowRef" json:"innerMessageFlowRef"`
	OuterMessageFlowRef MessageFlow `xml:"outerMessageFlowRef" json:"outerMessageFlowRef"`
}

func CreateMessageFlowAssociation(id string, innerMessageFlowRef MessageFlow, outerMessageFlowRef MessageFlow) MessageFlowAssociation {
	baseElement := foundation.CreateBaseElement(id)
	return MessageFlowAssociation{
		BaseElement:         baseElement,
		InnerMessageFlowRef: innerMessageFlowRef,
		OuterMessageFlowRef: outerMessageFlowRef,
	}
}

func (m MessageFlowAssociation) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, m, m.Id)
	checks = append(checks, m.BaseElement.Validate(name)...)
	checks = append(checks, m.InnerMessageFlowRef.Validate(name)...)
	checks = append(checks, m.OuterMessageFlowRef.Validate(name)...)
	return validation.FilterErrors(checks)
}
