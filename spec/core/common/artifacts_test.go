package common

import (
	"testing"

	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestArtifacts(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateArtifact("id"),
	)
	t.Run(name, fn)
}

func TestAssociation(t *testing.T) {
	ref1 := foundation.CreateBaseElement("id")
	ref2 := foundation.CreateBaseElement("id2")
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateAssociation("innerId", ref1, ref2, ADBoth),
	)
	t.Run(name, fn)
}

func TestGroup(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateGroup("id"),
	)
	t.Run(name, fn)
}

func TestCategory(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateCategory("id", "name"),
	)
	t.Run(name, fn)
}

func TestCategoryValue(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateCategoryValue("id", "value"),
	)
	t.Run(name, fn)
}

func TestTextAnnotation(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateTextAnnotation("id", "text"),
	)
	t.Run(name, fn)
}
