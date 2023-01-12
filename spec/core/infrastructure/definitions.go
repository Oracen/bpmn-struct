package infrastructure

import (
	"bpmn-struct/spec/core/foundation"
	"bpmn-struct/validation"
	"fmt"
)

type Definitions struct {
	Name               string                 `xml:"name" json:"name"`
	TargetNamespace    string                 `xml:"targetNamespace" json:"targetNamespace"`
	ExpressionLanguage []string               `xml:"expressionLanguage" json:"expressionLanguage"`
	TypeLanguage       []string               `xml:"typeLanguage" json:"typeLanguage"`
	RootElements       []RootElement          `xml:"rootElements" json:"rootElements"`
	Diagrams           []BPMNDiagram          `xml:"diagrams" json:"diagrams"`
	Imports            []Import               `xml:"imports" json:"imports"`
	Extensions         []foundation.Extension `xml:"extensions" json:"extensions"`
	Relationships      []Relationship         `xml:"relationships" json:"relationships"`
	Exporter           []string               `xml:"exporter" json:"exporter"`
	ExporterVersion    []string               `xml:"exporterVersion" json:"exporterVersion"`
}

func CreateDefinitions(name, targetNamespace string) Definitions {
	return Definitions{
		Name:               name,
		TargetNamespace:    targetNamespace,
		ExpressionLanguage: []string{"http://www.w3.org/1999/XPath"},
		TypeLanguage:       []string{"http://www.w3.org/2001/XMLSchema"},
		RootElements:       []RootElement{},
		Diagrams:           []BPMNDiagram{},
		Imports:            []Import{},
		Extensions:         []foundation.Extension{},
		Relationships:      []Relationship{},
		Exporter:           []string{},
		ExporterVersion:    []string{},
	}
}

func (d Definitions) Validate() (errors []error) {
	name := fmt.Sprintf("Definitions:%s", d.Name)
	checks := []error{
		validation.ValNonzero(name, "Name", d.Name),
		validation.ValNonzero(name, "TargetNamespace", d.TargetNamespace),
		validation.ArrZeroOne(name, "ExpressionLanguage", d.ExpressionLanguage),
		validation.ArrZeroOne(name, "TypeLanguage", d.TypeLanguage),
		validation.ArrZeroOne(name, "Exporter", d.Exporter),
		validation.ArrZeroOne(name, "ExporterVersion", d.ExporterVersion),
	}
	return validation.FilterErrors(checks)
}
