package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type CorrelationKey struct {
	foundation.BaseElement
	Name                   []string              `xml:"name" json:"name"`
	CorrelationPropertyRef []CorrelationProperty `xml:"correlationPropertyRef" json:"correlationPropertyRef"`
}

func CreateCorrelationKey(id string) CorrelationKey {
	baseElement := foundation.CreateBaseElement(id)
	return CorrelationKey{
		BaseElement:            baseElement,
		Name:                   []string{},
		CorrelationPropertyRef: []CorrelationProperty{},
	}
}

func (c CorrelationKey) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, c, c.Id)
	checks = append(checks, c.BaseElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, c.CorrelationPropertyRef)...)
	checks = append(checks, validation.ArrZeroOne(name, "Name", c.Name))
	return validation.FilterErrors(checks)
}

type CorrelationProperty struct {
	foundation.BaseElement
	Name                                   []string                                 `xml:"name" json:"name"`
	Type                                   []string                                 `xml:"type" json:"type"`
	CorrelationPropertyRetrievalExpression []CorrelationPropertyRetrievalExpression `xml:"correlationPropertyRetrievalExpression" json:"correlationPropertyRetrievalExpression"`
}

func CreateCorrelationProperty(id string, retrievalExpression CorrelationPropertyRetrievalExpression) CorrelationProperty {
	baseElement := foundation.CreateBaseElement(id)
	return CorrelationProperty{
		BaseElement:                            baseElement,
		Name:                                   []string{},
		Type:                                   []string{},
		CorrelationPropertyRetrievalExpression: []CorrelationPropertyRetrievalExpression{retrievalExpression},
	}
}

func (c CorrelationProperty) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, c, c.Id)
	checks = append(checks, c.BaseElement.Validate(name)...)
	checks = append(checks, validation.ArrOneOrMore(name, "CorrelationPropertyRetrievalExpression", c.CorrelationPropertyRetrievalExpression))
	checks = append(checks, validation.ArrCheckItems(name, c.CorrelationPropertyRetrievalExpression)...)
	return validation.FilterErrors(checks)
}

type CorrelationPropertyRetrievalExpression struct {
	foundation.BaseElement
	MessagePath FormalExpression `xml:"messagePath" json:"messagePath"`
	MessageRef  Message          `xml:"messageRef" json:"messageRef"`
}

func CreateCorrelationPropertyRetrievalExpression(id string, messagePath FormalExpression, messageRef Message) CorrelationPropertyRetrievalExpression {
	baseElement := foundation.CreateBaseElement(id)
	return CorrelationPropertyRetrievalExpression{
		BaseElement: baseElement,
		MessagePath: messagePath,
		MessageRef:  messageRef,
	}
}

func (c CorrelationPropertyRetrievalExpression) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, c, c.Id)
	checks = append(checks, c.BaseElement.Validate(name)...)
	checks = append(checks, c.MessagePath.Validate(name)...)
	checks = append(checks, c.MessageRef.Validate(name)...)

	return validation.FilterErrors(checks)
}

type CorrelationSubscription struct {
	foundation.BaseElement
	CorrelationKeyRef          CorrelationKey               `xml:"correlationKeyRef" json:"correlationKeyRef"`
	CorrelationPropertyBinding []CorrelationPropertyBinding `xml:"correlationPropertyBinding" json:"correlationPropertyBinding"`
}

func CreateCorrelationSubscription(id string, correlationKeyRef CorrelationKey) CorrelationSubscription {
	baseElement := foundation.CreateBaseElement(id)
	return CorrelationSubscription{
		BaseElement:                baseElement,
		CorrelationKeyRef:          correlationKeyRef,
		CorrelationPropertyBinding: []CorrelationPropertyBinding{},
	}
}

func (c CorrelationSubscription) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, c, c.Id)
	checks = append(checks, c.BaseElement.Validate(name)...)
	checks = append(checks, c.CorrelationKeyRef.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, c.CorrelationPropertyBinding)...)
	checks = append(checks, validation.ArrZeroOne(name, "CorrelationPropertyBinding", c.CorrelationPropertyBinding))
	return validation.FilterErrors(checks)
}

type CorrelationPropertyBinding struct {
	foundation.BaseElement
	DataPath               FormalExpression    `xml:"dataPath" json:"dataPath"`
	CorrelationPropertyRef CorrelationProperty `xml:"correlationPropertyRef" json:"correlationPropertyRef"`
}

func CreateCorrelationPropertyBinding(id string, dataPath FormalExpression, correlationPropertyRef CorrelationProperty) CorrelationPropertyBinding {
	baseElement := foundation.CreateBaseElement(id)
	return CorrelationPropertyBinding{
		BaseElement:            baseElement,
		DataPath:               dataPath,
		CorrelationPropertyRef: correlationPropertyRef,
	}
}

func (c CorrelationPropertyBinding) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, c, c.Id)
	checks = append(checks, c.BaseElement.Validate(name)...)
	checks = append(checks, c.DataPath.Validate(name)...)
	checks = append(checks, c.CorrelationPropertyRef.Validate(name)...)
	return validation.FilterErrors(checks)
}
