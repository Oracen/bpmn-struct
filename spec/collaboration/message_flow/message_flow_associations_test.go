package message_flow

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestMessageFlowAssociation(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateMessageFlowAssociation("id", CreateMessageFlow(
			"id2",
			CreateInteractionNode("id21"),
			CreateInteractionNode("id22"),
		),
			CreateMessageFlow(
				"id3",
				CreateInteractionNode("id31"),
				CreateInteractionNode("id32"),
			)),
	)
	t.Run(name, fn)
}
