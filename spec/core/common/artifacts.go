package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type AssociationDirection int

var (
	associationDirectionString = [...]string{"None", "One", "Both"}
)

const (
	ADNone AssociationDirection = iota
	ADOne
	ADBoth
)

func (r AssociationDirection) String() string {
	return associationDirectionString[r]
}

func (r AssociationDirection) EnumIndex() int {
	return int(r)
}

func (r AssociationDirection) ToEnum(input string) (relp AssociationDirection, err error) {
	for idx, item := range associationDirectionString {
		if item == input {
			return AssociationDirection(idx), nil
		}
	}
	err = validation.NewEnumValidationError("AssociationDirection", "String", associationDirectionString[:], input)
	return
}

type Artifact struct {
	foundation.BaseElement
}

func CreateArtifact(id string) Artifact {
	baseElement := foundation.CreateBaseElement(id)
	return Artifact{
		BaseElement: baseElement,
	}
}

func (a Artifact) Validate(name string) []error {
	return a.BaseElement.Validate(name)
}

type Association struct {
	Artifact
	AssociationDirection AssociationDirection   `xml:"associationDirection" json:"associationDirection"`
	SourceRef            foundation.BaseElement `xml:"sourceRef" json:"sourceRef"`
	TargetRef            foundation.BaseElement `xml:"targetRef" json:"targetRef"`
}

func CreateAssociation(id string, sourceRef, targetRef foundation.BaseElement, direction AssociationDirection) Association {
	return Association{
		Artifact:             CreateArtifact(id),
		AssociationDirection: direction,
		SourceRef:            sourceRef,
		TargetRef:            targetRef,
	}
}

func (a Association) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, a, a.Id)
	checks = append(checks, a.Artifact.Validate(name)...)
	checks = append(checks, a.SourceRef.Validate(name)...)
	checks = append(checks, a.TargetRef.Validate(name)...)

	return validation.FilterErrors(checks)
}

type Group struct {
	Artifact
	CategoryValueRef []CategoryValue `xml:"categoryValue" json:"categoryValue"`
}

func CreateGroup(id string) Group {
	return Group{
		Artifact:         CreateArtifact(id),
		CategoryValueRef: []CategoryValue{},
	}
}

func (g Group) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, g, g.Id)
	checks = append(checks, g.Artifact.Validate(name)...)
	checks = append(checks, validation.ArrZeroOne(name, "CategoryValueRef", g.CategoryValueRef))
	checks = append(checks, validation.ArrCheckItems(name, g.CategoryValueRef)...)

	return validation.FilterErrors(checks)
}

type Category struct {
	foundation.RootElement
	Name          string          `xml:"name" json:"name"`
	CategoryValue []CategoryValue `xml:"categoryValue" json:"categoryValue"`
}

func CreateCategory(id, name string) Category {
	rootElement := foundation.CreateRootElement(id)
	return Category{
		RootElement:   rootElement,
		Name:          name,
		CategoryValue: []CategoryValue{},
	}
}

func (c Category) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, c, c.Name)
	checks = append(checks, c.RootElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, c.CategoryValue)...)
	checks = append(checks, validation.ArrZeroOne(name, "CategoryValue", c.CategoryValue))
	checks = append(checks, validation.ValNonzero(name, "Name", c.Name))

	return validation.FilterErrors(checks)
}

type CategoryValue struct {
	foundation.BaseElement
	Value                   string        `xml:"value" json:"value"`
	Category                []Category    `xml:"category" json:"category"`
	CategorizedFlowElements []FlowElement `xml:"categorizedFlowElements" json:"categorizedFlowElements"`
}

func CreateCategoryValue(id, value string) CategoryValue {
	baseElement := foundation.CreateBaseElement(id)
	return CategoryValue{
		BaseElement:             baseElement,
		Value:                   value,
		Category:                []Category{},
		CategorizedFlowElements: []FlowElement{},
	}
}

func (c CategoryValue) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, c, c.Id)
	checks = append(checks, c.BaseElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, c.Category)...)
	checks = append(checks, validation.ArrCheckItems(name, c.CategorizedFlowElements)...)
	checks = append(checks, validation.ArrZeroOne(name, "Category", c.Category))
	checks = append(checks, validation.ValNonzero(name, "Value", c.Value))

	return validation.FilterErrors(checks)
}

type TextAnnotation struct {
	foundation.BaseElement
	Text       string `xml:"text" json:"text"`
	TextFormat string `xml:"textFormat" json:"textFormat"`
}

func CreateTextAnnotation(id, text string) TextAnnotation {
	baseElement := foundation.CreateBaseElement(id)
	return TextAnnotation{
		BaseElement: baseElement,
		Text:        text,
		TextFormat:  "text/plain",
	}
}

func (t TextAnnotation) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, t, t.Id)
	checks = append(checks, t.BaseElement.Validate(name)...)
	checks = append(checks, validation.ValNonzero(name, "Text", t.Text))
	checks = append(checks, validation.ValNonzero(name, "TextFormat", t.TextFormat))

	return validation.FilterErrors(checks)
}
