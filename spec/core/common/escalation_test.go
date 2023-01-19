package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestEscalation(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateEscalation("id", "name", "escalationCode"),
	)
	t.Run(name, fn)
}
