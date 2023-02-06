package pool_and_participant

import "testing"

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
		CreatePartnerEntity("id"),
	)
	t.Run(name, fn)
}

func TestPartnerRole(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreatePartnerRole("id"),
	)
	t.Run(name, fn)
}

func TestParticipantMultiplicity(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateParticipantMultiplicity("id"),
	)
	t.Run(name, fn)
}

func TestParticipantAssociation(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateParticipantAssociation("id"),
	)
	t.Run(name, fn)
}
