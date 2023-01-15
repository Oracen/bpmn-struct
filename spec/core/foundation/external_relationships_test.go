package foundation

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestRelationship(t *testing.T) {
	value := true
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateRelationship("id", "typeName", RDBoth, "any", value),
	)
	t.Run(name, fn)
}
