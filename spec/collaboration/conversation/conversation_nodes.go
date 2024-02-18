package conversation

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/common"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"

	"github.com/Oracen/bpmn-struct/spec/collaboration/message_flow"
	pap "github.com/Oracen/bpmn-struct/spec/collaboration/pool_and_participant"
)

type ConversationNode struct {
	foundation.BaseElement
	Name            []string                   `xml:"name" json:"name"`
	ParticipantRefs []pap.Participant          `xml:"participantRefs" json:"participantRefs"`
	MessageFlowRefs []message_flow.MessageFlow `xml:"messageFlowRefs" json:"messageFlowRefs"`
	CorrelationKeys []common.CorrelationKey    `xml:"correlationKeys" json:"correlationKeys"`
}

func CreateConversationNode(id string) ConversationNode {
	baseElement := foundation.CreateBaseElement(id)
	return ConversationNode{
		BaseElement:     baseElement,
		Name:            []string{},
		ParticipantRefs: []pap.Participant{},
		MessageFlowRefs: []message_flow.MessageFlow{},
		CorrelationKeys: []common.CorrelationKey{},
	}
}

func (c ConversationNode) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, c, c.Id)
	checks = append(checks, c.BaseElement.Validate(name)...)
	checks = append(checks, validation.ArrZeroOne(name, "Name", c.Name))
	checks = append(checks, validation.ArrTwoOrMore(name, "ParticipantRefs", c.ParticipantRefs))
	return validation.FilterErrors(checks)
}
