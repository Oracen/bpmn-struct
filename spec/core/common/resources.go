package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type Resource struct {
	foundation.RootElement
	Name               string              `xml:"name" json:"name"`
	ResourceParameters []ResourceParameter `xml:"resourceParameters" json:"resourceParameters"`
}

func CreateResource(id, name string) Resource {
	rootElement := foundation.CreateRootElement(id)
	return Resource{
		RootElement:        rootElement,
		Name:               name,
		ResourceParameters: []ResourceParameter{},
	}
}

func (r Resource) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, r, r.Id)
	checks = append(checks, r.RootElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, r.ResourceParameters)...)
	return validation.FilterErrors(checks)
}

type ResourceParameter struct {
	foundation.RootElement
	Name       string           `xml:"name" json:"name"`
	ItemRef    []ItemDefinition `xml:"itemRef" json:"itemRef"`
	IsRequired bool             `xml:"isRequired" json:"isRequired"`
}

func CreateResourceParameter(id, name string) ResourceParameter {
	rootElement := foundation.CreateRootElement(id)
	return ResourceParameter{
		RootElement: rootElement,
		Name:        name,
		ItemRef:     []ItemDefinition{},
		IsRequired:  false,
	}
}

func (r ResourceParameter) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, r, r.Id)
	checks = append(checks, r.RootElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, r.ItemRef)...)

	return validation.FilterErrors(checks)
}
