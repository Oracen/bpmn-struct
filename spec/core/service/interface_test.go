package service

import (
	"testing"

	"github.com/Oracen/bpmn-struct/spec/core/common"
	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestInterface(t *testing.T) {
	operations := []Operation{CreateOperation("id", "text", common.CreateMessage("id"))}
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateInterface("id", "name", operations),
	)
	t.Run(name, fn)
}
