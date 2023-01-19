package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestItemDefinition(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateItemDefinition("id", ITEM_KIND_Physical),
	)
	t.Run(name, fn)
}
