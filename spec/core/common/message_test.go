package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestMessage(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateMessage("id"),
	)
	t.Run(name, fn)
}
