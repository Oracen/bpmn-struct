package infrastructure

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestImport(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateImport("id", "namespace"),
	)
	t.Run(name, fn)
}
