package service

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type EndPoint struct {
	foundation.RootElement
}

func CreateEndPoint(id string) EndPoint {
	rootElement := foundation.CreateRootElement(id)
	return EndPoint{
		RootElement: rootElement,
	}
}

func (e EndPoint) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, e, e.Id)
	checks = append(checks, e.RootElement.Validate(name)...)
	return validation.FilterErrors(checks)
}
