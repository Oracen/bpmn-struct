package foundation

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestExtension(t *testing.T) {
	testhelpers.CreateDefaultIsValid(
		t,
		CreateExtension("name"),
	)
}

func TestExtensionDefinition(t *testing.T) {
	testhelpers.CreateDefaultIsValid(
		t,
		CreateExtensionDefinition("name"),
	)
}

func TestExtensionAttributeDefinition(t *testing.T) {
	testhelpers.CreateDefaultIsValid(
		t,
		CreateExtensionAttributeDefinition("name", "typeName"),
	)
}

func TestExtensionAttributeValue(t *testing.T) {
	testhelpers.CreateDefaultIsValid(
		t,
		CreateExtensionAttributeValue("name", "typeName", "anyvalue", false),
	)
}
