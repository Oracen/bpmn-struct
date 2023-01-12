package foundation

import (
	"bpmn-struct/validation"
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type Extension struct {
	MustUnderstand []bool              `xml:"mustUnderstand" json:"mustUnderstand"`
	Definition     ExtensionDefinition `xml:"definition" json:"definition"`
}

func CreateExtension(definition ExtensionDefinition) Extension {
	return Extension{
		MustUnderstand: []bool{false},
		Definition:     definition,
	}
}

func (e Extension) Validate() (errors []error) {
	json, _ := json.Marshal(e)
	name := fmt.Sprintf("Extension:%s", md5.Sum(json))
	checks := []error{
		validation.ArrZeroOne(name, "MustUnderstand", e.MustUnderstand),
		validation.ValNonzero(name, "Definition", e.Definition),
	}
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

func (e ExtensionDefinition) Validate() (errors []error) {
	name := fmt.Sprintf("ExtensionDefinition:%s", e.Name)
	checks := []error{
		validation.ValNonzero(name, "Name", e.Name),
	}
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

func (e ExtensionAttributeDefinition) Validate() (errors []error) {
	name := fmt.Sprintf("ExtensionAttributeDefinition:%s", e.Name)
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

func CreateExtensionAttributeValue(definition ExtensionAttributeDefinition, value any, isRef bool) ExtensionAttributeValue {
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

func (e ExtensionAttributeValue) Validate() (errors []error) {
	value := append(e.Value, e.ValueRef...)[0]
	name := fmt.Sprintf("ExtensionAttributeValue:%s", value)
	checks := []error{
		validation.ArrZeroOne(name, "Value", e.Value),
		validation.ArrZeroOne(name, "ValueRef", e.ValueRef),
		validation.ArraysMaxCount(name, "Value,ValueRef", 1, e.Value, e.ValueRef),
		validation.ValNonzero(name, "Type", e.ExtensionAttributeDefinition),
	}
	return validation.FilterErrors(checks)
}
