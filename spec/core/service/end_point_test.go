package service

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestEndPoint(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateEndPoint("id"),
	)
	t.Run(name, fn)
}
