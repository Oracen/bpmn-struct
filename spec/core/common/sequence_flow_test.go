package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestSequenceFlow(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateSequenceFlow("id", "source", "target"),
	)
	t.Run(name, fn)
}
