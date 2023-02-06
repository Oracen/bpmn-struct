package collaboration

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestCollaboration(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateCollaboration("id", "name"),
	)
	t.Run(name, fn)
}
