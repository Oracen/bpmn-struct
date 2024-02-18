package collaboration

import (
	"github.com/Oracen/bpmn-struct/shared"
	messageflow "github.com/Oracen/bpmn-struct/spec/collaboration/message_flow"
	pap "github.com/Oracen/bpmn-struct/spec/collaboration/pool_and_participant"
	"github.com/Oracen/bpmn-struct/spec/core/common"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type Collaboration struct {
	foundation.RootElement
	Name                     string                               `xml:"name" json:"name"`
	ChoreographyRef          []Choreography                       `xml:"choreographyRef" json:"choreographyRef"`
	CorrelationKeys          []common.CorrelationKey              `xml:"correlationKeys" json:"correlationKeys"`
	ConversationAssociations []ConversationAssociation            `xml:"conversationAssociations" json:"conversationAssociations"`
	Conversations            []ConversationNode                   `xml:"conversations" json:"conversations"`
	ConversationLinks        []ConversationLink                   `xml:"conversationLinks" json:"conversationLinks"`
	Artifacts                []common.Artifact                    `xml:"artifacts" json:"artifacts"`
	Participants             []pap.Participant                    `xml:"participants" json:"participants"`
	ParticipantAssociations  []pap.ParticipantAssociation         `xml:"participantAssociations" json:"participantAssociations"`
	MessageFlow              []messageflow.MessageFlow            `xml:"messageFlow" json:"messageFlow"`
	MessageFlowAssociations  []messageflow.MessageFlowAssociation `xml:"messageFlowAssociations" json:"messageFlowAssociations"`
	IsClosed                 bool                                 `xml:"isClosed" json:"isClosed"`
}

func CreateCollaboration(id, name string) Collaboration {
	rootElement := foundation.CreateRootElement(id)
	return Collaboration{
		RootElement: rootElement,
		Name:        name,
	}
}

func (c Collaboration) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, c, c.Id)
	checks = append(checks, c.RootElement.Validate(name)...)
	return validation.FilterErrors(checks)
}
