package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestEvents(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateEvent("id"),
	)
	t.Run(name, fn)
}
