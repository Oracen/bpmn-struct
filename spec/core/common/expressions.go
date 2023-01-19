package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type Expression struct {
	foundation.BaseElement
}

func CreateExpression(id string) Expression {
	baseElement := foundation.CreateBaseElement(id)
	return Expression{
		BaseElement: baseElement,
	}
}

func (e Expression) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, e, e.Id)
	checks = append(checks, e.BaseElement.Validate(name)...)
	return validation.FilterErrors(checks)
}

type FormalExpression struct {
	Expression
	Language           []string       `xml:"language" json:"language"`
	Body               any            `xml:"body" json:"body"`
	EvaluatesToTypeRef ItemDefinition `xml:"evaluatesToTypeRef" json:"evaluatesToTypeRef"`
}

func CreateFormalExpression(id string, body any, evaluatesToTypeRef ItemDefinition) FormalExpression {
	expression := CreateExpression(id)
	return FormalExpression{
		Expression:         expression,
		Language:           []string{},
		Body:               body,
		EvaluatesToTypeRef: evaluatesToTypeRef,
	}
}

func (f FormalExpression) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, f, f.Id)
	checks = append(checks, f.Expression.Validate(name)...)
	checks = append(checks, f.EvaluatesToTypeRef.Validate(name)...)
	checks = append(
		checks,
		validation.ArrZeroOne(name, "Language", f.Language),
		validation.ValNonzero(name, "Body", f.Body),
	)
	return validation.FilterErrors(checks)
}
