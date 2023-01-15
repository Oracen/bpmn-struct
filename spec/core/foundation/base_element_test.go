package foundation

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestBaseElement(t *testing.T) {
	testhelpers.CreateDefaultIsValid(
		t,
		CreateBaseElement("id"),
	)
}
