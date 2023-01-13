package foundation

import (
	"bpmn-struct/constants"
	"bpmn-struct/validation"
	"fmt"
)

type Documentation struct {
	BaseElement
	Text       string `xml:"text" json:"text"`
	TextFormat string `xml:"textFormat" json:"textFormat"`
}

func CreateDocumentation(id, text string) Documentation {
	baseElement := CreateBaseElement(id)
	return Documentation{
		BaseElement: baseElement,
		Text:        text,
		TextFormat:  "text/plain",
	}
}

func (d Documentation) Validate(name string) (errors []error) {
	if name == "" {
		name = fmt.Sprintf("Import:%s", d.Text[:constants.IdentifierTruncationLen])
	}

	checksBase := d.BaseElement.Validate(name)
	checksCurrent := []error{
		validation.ValNonzero(name, "ImportType", d.Text),
		validation.ValNonzero(name, "Namespace", d.TextFormat),
	}
	return validation.FilterErrors(append(checksBase, checksCurrent...))
}
