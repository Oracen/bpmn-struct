package foundation

import (
	"bpmn-struct/validation"
	"fmt"
)

type RelationshipDirection int

var (
	relationshipDirectionString = [...]string{"None", "Forward", "Backward", "Both"}
)

const (
	RDNone RelationshipDirection = iota
	RDForward
	RDBackward
	RDBoth
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

func CreateRelationship(id, typeName string, direction RelationshipDirection, sources, targets []any) Relationship {
	baseElement := CreateBaseElement(id)
	return Relationship{
		BaseElement: baseElement,
		Type:        typeName,
		Direction:   direction,
		Sources:     sources,
		Targets:     targets,
	}
}

func (r Relationship) Validate(name string) []error {
	if name == "" {
		name = fmt.Sprintf("ExtensionAttributeValue:%s", r.Type)
	}
	checksBase := r.BaseElement.Validate(name)
	checksCurrent := []error{
		validation.ValNonzero(name, "Type", r.Type),
		validation.ArrOneOrMore(name, "Sources", r.Sources),
		validation.ArrOneOrMore(name, "Targets", r.Targets),
	}
	return validation.FilterErrors(append(checksBase, checksCurrent...))
}
