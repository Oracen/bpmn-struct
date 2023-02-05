package foundation

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestRelationship(t *testing.T) {
	value := true
	sources := []any{"any"}
	targets := []any{value}
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateRelationship("id", "typeName", RELATIONSHIP_DIRECTION_Both, sources, targets),
	)
	t.Run(name, fn)
}
