package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestFlowElements(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateFlowElement("id"),
	)
	t.Run(name, fn)
}

func TestFlowNode(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateFlowNode("id"),
	)
	t.Run(name, fn)
}
