package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestGateways(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateGateway("id"),
	)
	t.Run(name, fn)
}
