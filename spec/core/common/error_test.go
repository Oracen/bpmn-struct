package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestError(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateError("id", "name", "errorCode"),
	)
	t.Run(name, fn)
}
