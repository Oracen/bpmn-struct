package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestResources(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateResource("id", "name"),
	)
	t.Run(name, fn)
}

func TestResourceDefinition(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateResourceParameter("id", "name"),
	)
	t.Run(name, fn)
}
