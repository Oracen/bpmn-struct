package pool_and_participant

import (
	"fmt"

	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/spec/core/service"
	"github.com/Oracen/bpmn-struct/spec/process"
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

func (p Participant) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, p, p.Id)
	checks = append(checks, p.BaseElement.Validate(name)...)
	checks = append(checks, validation.ArrNonzero(name, "Name", p.Name)...)
	checks = append(checks, validation.ArrCheckItems(name, p.ProcessRef)...)
	checks = append(checks, validation.ArrCheckItems(name, p.PartnerRoleRef)...)
	checks = append(checks, validation.ArrCheckItems(name, p.PartnerEntityRef)...)
	checks = append(checks, validation.ArrCheckItems(name, p.InterfaceRef)...)
	checks = append(checks, validation.ArrCheckItems(name, p.ParticipantMultiplicity)...)
	checks = append(checks, validation.ArrCheckItems(name, p.EndPointRefs)...)
	checks = append(
		checks,
		validation.ArrZeroOne(name, "Name", p.Name),
		validation.ArrZeroOne(name, "ProcessRef", p.ProcessRef),
		validation.ArrZeroOne(name, "ParticipantMultiplicity", p.ParticipantMultiplicity),
		validation.ArrZeroOne(name, "EndPointRefs", p.EndPointRefs),
	)
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

func (p PartnerEntity) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, p, p.Id)
	checks = append(checks, p.BaseElement.Validate(name)...)
	checks = append(checks, p.ParticipantRef.Validate(name)...)
	checks = append(checks, validation.ValNonzero(name, "Name", p.Name))
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

func (p PartnerRole) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, p, p.Id)
	checks = append(checks, p.BaseElement.Validate(name)...)
	checks = append(checks, p.ParticipantRef.Validate(name)...)
	checks = append(checks, validation.ValNonzero(name, "Name", p.Name))
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

func (p ParticipantMultiplicity) Validate(name string) []error {
	checks := []error{}

	identifier := fmt.Sprintf("%v-%v-%v", p.Minimum, p.Maximum, p.NumParticipants)
	name = shared.TypeNameString(name, p, identifier)
	checks = append(checks, validation.ArrGTEZero(name, "Maximum", p.Maximum)...)
	checks = append(checks, validation.ArrGTEZero(name, "NumParticipants", p.NumParticipants)...)

	checks = append(
		checks,
		validation.ValGTEZero(name, "Minimum", p.Minimum),
		validation.ArrZeroOne(name, "Maximum", p.Maximum),
		validation.ArrZeroOne(name, "NumParticipants", p.NumParticipants),
	)
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

func (p ParticipantAssociation) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, p, p.Id)
	checks = append(checks, p.BaseElement.Validate(name)...)
	checks = append(checks, p.InnerParticipantRef.Validate(name)...)
	checks = append(checks, p.OuterParticipantRef.Validate(name)...)
	return validation.FilterErrors(checks)
}
