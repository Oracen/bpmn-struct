package testhelpers

import (
	"fmt"
	"testing"

	"github.com/Oracen/bpmn-struct/shared"
	"github.com/stretchr/testify/assert"
)

func CreateDefaultIsValid(t *testing.T, testStruct shared.BPMNStruct) (string, func(*testing.T)) {
	name := fmt.Sprintf("default %T struct is valid on creation", testStruct)
	fn := func(t *testing.T) {
		t.Helper()
		errors := testStruct.Validate("")
		assert.Len(t, errors, 0)
	}
	return name, fn

}
