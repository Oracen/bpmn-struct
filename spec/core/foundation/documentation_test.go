package foundation

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestDocumentation(t *testing.T) {
	testhelpers.CreateDefaultIsValid(
		t,
		CreateDocumentation("id", "a docstring"),
	)
}
