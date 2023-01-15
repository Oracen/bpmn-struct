package foundation

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestExtension(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateExtension("name"),
	)
	t.Run(name, fn)
}

func TestExtensionDefinition(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateExtensionDefinition("name"),
	)
	t.Run(name, fn)
}

func TestExtensionAttributeDefinition(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateExtensionAttributeDefinition("name", "typeName"),
	)
	t.Run(name, fn)
}

func TestExtensionAttributeValue(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateExtensionAttributeValue("name", "typeName", "anyvalue", false),
	)
	t.Run(name, fn)
}
