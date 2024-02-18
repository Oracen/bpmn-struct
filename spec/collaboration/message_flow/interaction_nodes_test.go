package message_flow

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestInteractionNode(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateInteractionNode(),
	)
	t.Run(name, fn)
}
