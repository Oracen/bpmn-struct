package pool_and_participant

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/spec/core/service"
	"github.com/Oracen/bpmn-struct/validation"
)

type Participant struct {
	foundation.BaseElement
	Name                    []string                  `xml:"name" json:"name"`
	ProcessRef              []process.Process         `xml:"processRef" json:"processRef"`
	PartnerRoleRef          []PartnerRole             `xml:"partnerRoleRef" json:"partnerRoleRef"`
	PartnerEntityRef        []PartnerEntity           `xml:"partnerEntityRef" json:"partnerEntityRef"`
	InterfaceRef            []service.Interface       `xml:"interfaceRef" json:"interfaceRef"`
	ParticipantMultiplicity []ParticipantMultiplicity `xml:"participantMultiplicity" json:"participantMultiplicity"`
	EndPointRefs            []service.EndPoint        `xml:"endPointRefs" json:"endPointRefs"`
}

func CreateParticipant(id string) Participant {
	baseElement := foundation.CreateBaseElement(id)
	return Participant{
		BaseElement:             baseElement,
		Name:                    []string{},
		ProcessRef:              []process.Process{},
		PartnerRoleRef:          []PartnerRole{},
		PartnerEntityRef:        []PartnerEntity{},
		InterfaceRef:            []service.Interface{},
		ParticipantMultiplicity: []ParticipantMultiplicity{},
		EndPointRefs:            []service.EndPoint{},
	}
}

func (x Participant) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, t, t.Id)
	checks = append(checks, t.BaseElement.Validate(name)...)
	return validation.FilterErrors(checks)
}

type PartnerEntity struct {
	foundation.BaseElement
	Name           string      `xml:"name" json:"name"`
	ParticipantRef Participant `xml:"participantRef" json:"participantRef"`
}

func CreatePartnerEntity(id, name string, participant Participant) PartnerEntity {
	baseElement := foundation.CreateBaseElement(id)
	return PartnerEntity{
		BaseElement:    baseElement,
		Name:           name,
		ParticipantRef: participant,
	}
}

func (x PartnerEntity) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, t, t.Id)
	checks = append(checks, t.BaseElement.Validate(name)...)
	return validation.FilterErrors(checks)
}

type PartnerRole struct {
	foundation.BaseElement
	Name           string      `xml:"name" json:"name"`
	ParticipantRef Participant `xml:"participantRef" json:"participantRef"`
}

func CreatePartnerRole(id, name string, participant Participant) PartnerRole {
	baseElement := foundation.CreateBaseElement(id)
	return PartnerRole{
		BaseElement:    baseElement,
		Name:           name,
		ParticipantRef: participant,
	}
}

func (x PartnerRole) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, t, t.Id)
	checks = append(checks, t.BaseElement.Validate(name)...)
	return validation.FilterErrors(checks)
}

type ParticipantMultiplicity struct {
	Minimum         int   `xml:"minimum" json:"minimum"`
	Maximum         []int `xml:"maximum" json:"maximum"`
	NumParticipants []int `xml:"numParticipants" json:"numParticipants"`
}

func CreateParticipantMultiplicity() ParticipantMultiplicity {
	return ParticipantMultiplicity{
		Minimum:         0,
		Maximum:         []int{},
		NumParticipants: []int{},
	}
}

func (x ParticipantMultiplicity) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, t, t.Id)
	checks = append(checks, t.BaseElement.Validate(name)...)
	return validation.FilterErrors(checks)
}

type ParticipantAssociation struct {
	foundation.BaseElement
	InnerParticipantRef Participant `xml:"innerParticipantRef" json:"innerParticipantRef"`
	OuterParticipantRef Participant `xml:"outerParticipantRef" json:"outerParticipantRef"`
}

func CreateParticipantAssociation(id string, innerParticipant, outerParticipant Participant) ParticipantAssociation {
	baseElement := foundation.CreateBaseElement(id)
	return ParticipantAssociation{
		BaseElement:         baseElement,
		InnerParticipantRef: innerParticipant,
		OuterParticipantRef: outerParticipant,
	}
}

func (x ParticipantAssociation) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, t, t.Id)
	checks = append(checks, t.BaseElement.Validate(name)...)
	return validation.FilterErrors(checks)
}
