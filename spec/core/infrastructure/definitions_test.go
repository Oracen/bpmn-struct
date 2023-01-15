package infrastructure

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestDefinitions(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateDefinitions("id", "name", "targetNamespace"),
	)
	t.Run(name, fn)
}
