package pool_and_participant

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

func TestParticipant(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateParticipant("id"),
	)
	t.Run(name, fn)
}

func TestPartnerEntity(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreatePartnerEntity("id", "name", CreateParticipant("id")),
	)
	t.Run(name, fn)
}

func TestPartnerRole(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreatePartnerRole("id", "name", CreateParticipant("id")),
	)
	t.Run(name, fn)
}

func TestParticipantMultiplicity(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateParticipantMultiplicity(),
	)
	t.Run(name, fn)
}

func TestParticipantAssociation(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateParticipantAssociation("id", CreateParticipant("id2"), CreateParticipant("id3")),
	)
	t.Run(name, fn)
}
