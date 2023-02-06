package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestExpressions(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateFormalExpression("id", true, CreateItemDefinition("id", ITEM_KIND_Information)),
	)
	t.Run(name, fn)
}
