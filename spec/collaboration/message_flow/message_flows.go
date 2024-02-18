package message_flow

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/common"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type MessageFlowRef string

type MessageFlow struct {
	foundation.BaseElement
	SourceRef  InteractionNode  `xml:"sourceRef" json:"sourceRef"`
	TargetRef  InteractionNode  `xml:"targetRef" json:"targetRef"`
	MessageRef []common.Message `xml:"messageRef" json:"messageRef"`
}

func CreateMessageFlow(id string, sourceRef InteractionNode, targetRef InteractionNode) MessageFlow {
	baseElement := foundation.CreateBaseElement(id)
	return MessageFlow{
		BaseElement: baseElement,
		SourceRef:   sourceRef,
		TargetRef:   targetRef,
		MessageRef:  []common.Message{},
	}
}

func (m MessageFlow) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, m, m.Id)
	checks = append(checks, m.BaseElement.Validate(name)...)
	checks = append(checks, m.SourceRef.Validate(name)...)
	checks = append(checks, m.TargetRef.Validate(name)...)
	checks = append(checks, validation.ArrZeroOne(name, "MessageRef", m.MessageRef))
	return validation.FilterErrors(checks)
}
