package foundation

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestRootElement(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateRootElement("id"),
	)
	t.Run(name, fn)
}
