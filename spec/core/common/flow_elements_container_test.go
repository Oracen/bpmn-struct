package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestFlowElementsContainer(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateFlowElementsContainer("id"),
	)
	t.Run(name, fn)
}
