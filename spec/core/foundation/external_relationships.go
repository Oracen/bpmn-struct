package foundation

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/validation"
)

type RelationshipDirection int

var (
	relationshipDirectionString = [...]string{"None", "Forward", "Backward", "Both"}
)

const (
	RELATIONSHIP_DIRECTION_None RelationshipDirection = iota
	RELATIONSHIP_DIRECTION_Forward
	RELATIONSHIP_DIRECTION_Backward
	RELATIONSHIP_DIRECTION_Both
)

func (r RelationshipDirection) String() string {
	return relationshipDirectionString[r]
}

func (r RelationshipDirection) EnumIndex() int {
	return int(r)
}

func (r RelationshipDirection) ToEnum(input string) (relp RelationshipDirection, err error) {
	for idx, item := range relationshipDirectionString {
		if item == input {
			return RelationshipDirection(idx), nil
		}
	}
	err = validation.NewEnumValidationError("RelationshipDirection", "String", relationshipDirectionString[:], input)
	return
}

type Relationship struct {
	BaseElement
	Type      string                `xml:"type" json:"type"`
	Direction RelationshipDirection `xml:"relationshipDirection" json:"relationshipDirection"`
	Sources   []any                 `xml:"sources" json:"sources"`
	Targets   []any                 `xml:"targets" json:"targets"`
}

func CreateRelationship(id, typeName string, direction RelationshipDirection, source, target any) Relationship {
	baseElement := CreateBaseElement(id)
	return Relationship{
		BaseElement: baseElement,
		Type:        typeName,
		Direction:   direction,
		Sources:     []any{source},
		Targets:     []any{target},
	}
}

func (r Relationship) Validate(name string) []error {
	checks := []error{}
	name = shared.TypeNameString(name, r, r.Type)
	checks = append(checks, r.BaseElement.Validate(name)...)
	checks = append(checks,
		validation.ValNonzero(name, "Type", r.Type),
		validation.ArrOneOrMore(name, "Sources", r.Sources),
		validation.ArrOneOrMore(name, "Targets", r.Targets),
	)
	return validation.FilterErrors(checks)
}
