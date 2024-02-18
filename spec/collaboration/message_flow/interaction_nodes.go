package message_flow

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/collaboration/conversation"
	"github.com/Oracen/bpmn-struct/spec/collaboration/pool_and_participant"
	"github.com/Oracen/bpmn-struct/spec/process/activity"
	"github.com/Oracen/bpmn-struct/spec/process/event"
	"github.com/Oracen/bpmn-struct/validation"
)

type InteractionNode struct {
	pool_and_participant.Participant
	conversation.ConversationNode
	activity.Task
	event.Event
}

func CreateInteractionNode(id string) InteractionNode {
	participant := pool_and_participant.CreateParticipant(id)
	conversationNode := conversation.CreateConversationNode()
	task := activity.CreateTask()
	event := event.CreateEvent(id)
	return InteractionNode{
		Participant:      participant,
		ConversationNode: conversationNode,
		Task:             task,
		Event:            event,
	}
}

func (i InteractionNode) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, i, i.Id)
	checks = append(checks, i.Participant.Validate(name)...)
	checks = append(checks, i.ConversationNode.Validate(name)...)
	checks = append(checks, i.Task.Validate(name)...)
	checks = append(checks, i.Event.Validate(name)...)
	return validation.FilterErrors(checks)
}
