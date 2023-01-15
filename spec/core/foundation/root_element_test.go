package foundation

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestRootElement(t *testing.T) {
	testhelpers.CreateDefaultIsValid(
		t,
		CreateRootElement("id"),
	)
}
