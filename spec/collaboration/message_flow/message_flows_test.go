package message_flow

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestMessageFlow(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateMessageFlow("id", CreateInteractionNode("id"), CreateInteractionNode("id")),
	)
	t.Run(name, fn)
}
