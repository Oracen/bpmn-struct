package service

import (
	"testing"

	"github.com/Oracen/bpmn-struct/spec/core/common"
	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestOperation(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateOperation("id", "name", common.CreateMessage("id")),
	)
	t.Run(name, fn)
}
