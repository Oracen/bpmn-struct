package infrastructure

import (
	"fmt"

	bpmndi "github.com/Oracen/bpmn-struct/spec/bpmn_di"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/validation"
)

type Definitions struct {
	foundation.BaseElement
	Name               string                    `xml:"name" json:"name"`
	TargetNamespace    string                    `xml:"targetNamespace" json:"targetNamespace"`
	ExpressionLanguage []string                  `xml:"expressionLanguage" json:"expressionLanguage"`
	TypeLanguage       []string                  `xml:"typeLanguage" json:"typeLanguage"`
	RootElements       []foundation.RootElement  `xml:"rootElements" json:"rootElements"`
	Diagrams           []bpmndi.BPMNDiagram      `xml:"diagrams" json:"diagrams"`
	Imports            []Import                  `xml:"imports" json:"imports"`
	Extensions         []foundation.Extension    `xml:"extensions" json:"extensions"`
	Relationships      []foundation.Relationship `xml:"relationships" json:"relationships"`
	Exporter           []string                  `xml:"exporter" json:"exporter"`
	ExporterVersion    []string                  `xml:"exporterVersion" json:"exporterVersion"`
}

func CreateDefinitions(id, name, targetNamespace string) Definitions {
	baseElement := foundation.CreateBaseElement(id)
	return Definitions{
		BaseElement:        baseElement,
		Name:               name,
		TargetNamespace:    targetNamespace,
		ExpressionLanguage: []string{"http://www.w3.org/1999/XPath"},
		TypeLanguage:       []string{"http://www.w3.org/2001/XMLSchema"},
		RootElements:       []foundation.RootElement{},
		Diagrams:           []bpmndi.BPMNDiagram{},
		Imports:            []Import{},
		Extensions:         []foundation.Extension{},
		Relationships:      []foundation.Relationship{},
		Exporter:           []string{},
		ExporterVersion:    []string{},
	}
}

func (d Definitions) Validate(name string) (errors []error) {
	if name == "" {
		name = fmt.Sprintf("Definitions:%s", d.Name)
	}

	checksBase := d.BaseElement.Validate(name)
	checksCurrent := []error{
		validation.ValNonzero(name, "Name", d.Name),
		validation.ValNonzero(name, "TargetNamespace", d.TargetNamespace),
		validation.ArrZeroOne(name, "ExpressionLanguage", d.ExpressionLanguage),
		validation.ArrZeroOne(name, "TypeLanguage", d.TypeLanguage),
		validation.ArrZeroOne(name, "Exporter", d.Exporter),
		validation.ArrZeroOne(name, "ExporterVersion", d.ExporterVersion),
	}
	return validation.FilterErrors(append(checksBase, checksCurrent...))
}
