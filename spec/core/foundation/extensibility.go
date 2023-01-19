package foundation

import (
	"fmt"

	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/validation"
)

type Extension struct {
	MustUnderstand []bool              `xml:"mustUnderstand" json:"mustUnderstand"`
	Definition     ExtensionDefinition `xml:"definition" json:"definition"`
}

func CreateExtension(name string) Extension {
	definition := CreateExtensionDefinition(name)
	return Extension{
		MustUnderstand: []bool{false},
		Definition:     definition,
	}
}

func (e Extension) Validate(name string) (errors []error) {
	checks := []error{}
	name = shared.TypeNameString(name, e, shared.HashMd5(e))
	checks = append(checks, e.Definition.Validate(name)...)
	checks = append(
		checks,
		validation.ArrZeroOne(name, "MustUnderstand", e.MustUnderstand),
		validation.ValNonzero(name, "Definition", e.Definition),
	)
	return validation.FilterErrors(checks)
}

type ExtensionDefinition struct {
	Name                          string                         `xml:"name" json:"name"`
	ExtensionAttributeDefinitions []ExtensionAttributeDefinition `xml:"extensionAttributeDefinitions" json:"extensionAttributeDefinitions"`
}

func CreateExtensionDefinition(name string) ExtensionDefinition {
	return ExtensionDefinition{
		Name:                          name,
		ExtensionAttributeDefinitions: []ExtensionAttributeDefinition{},
	}
}

func (e ExtensionDefinition) Validate(name string) (errors []error) {
	checks := []error{}
	name = shared.TypeNameString(name, e, e.Name)
	checks = append(checks, validation.ArrCheckItems(name, e.ExtensionAttributeDefinitions)...)
	checks = append(checks, validation.ValNonzero(name, "Name", e.Name))
	return validation.FilterErrors(checks)
}

type ExtensionAttributeDefinition struct {
	Name        string `xml:"name" json:"name"`
	Type        string `xml:"type" json:"type"`
	IsReference []bool `xml:"isReference" json:"isReference"`
}

func CreateExtensionAttributeDefinition(name, typeName string) ExtensionAttributeDefinition {
	return ExtensionAttributeDefinition{
		Name:        name,
		Type:        typeName,
		IsReference: []bool{false},
	}
}

func (e ExtensionAttributeDefinition) Validate(name string) (errors []error) {
	name = shared.TypeNameString(name, e, e.Name)
	checks := []error{
		validation.ValNonzero(name, "Name", e.Name),
		validation.ValNonzero(name, "Type", e.Name),
		validation.ArrZeroOne(name, "IsReference", e.IsReference),
	}
	return validation.FilterErrors(checks)
}

type ExtensionAttributeValue struct {
	Value                        []any                        `xml:"value" json:"value"`
	ValueRef                     []any                        `xml:"valueRef" json:"valueRef"`
	ExtensionAttributeDefinition ExtensionAttributeDefinition `xml:"extensionAttributeDefinition" json:"extensionAttributeDefinition"`
}

func CreateExtensionAttributeValue(name, typeName string, value any, isRef bool) ExtensionAttributeValue {
	definition := CreateExtensionAttributeDefinition(name, typeName)
	valueItem, valueRef := []any{}, []any{}
	if isRef {
		valueRef = append(valueRef, value)
	} else {
		valueItem = append(valueItem, value)
	}
	return ExtensionAttributeValue{
		Value:                        valueItem,
		ValueRef:                     valueRef,
		ExtensionAttributeDefinition: definition,
	}
}

func (e ExtensionAttributeValue) Validate(name string) (errors []error) {
	checks := []error{}
	value := fmt.Sprintf("%s", append(e.Value, e.ValueRef...)[0])
	name = shared.TypeNameString(name, e, value)
	checks = append(checks, e.ExtensionAttributeDefinition.Validate(name)...)
	checks = append(
		checks,
		validation.ArrZeroOne(name, "Value", e.Value),
		validation.ArrZeroOne(name, "ValueRef", e.ValueRef),
		validation.ArraysMaxCount(name, "Value,ValueRef", 1, e.Value, e.ValueRef),
		validation.ValNonzero(name, "ExtensionAttributeDefinition", e.ExtensionAttributeDefinition),
	)
	return validation.FilterErrors(checks)
}
