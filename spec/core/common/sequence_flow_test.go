package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestSequenceFlow(t *testing.T) {
	sourceRef := CreateFlowNode("source")
	targetRef := CreateFlowNode("target")
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateSequenceFlow("id", sourceRef, targetRef),
	)
	t.Run(name, fn)
}
