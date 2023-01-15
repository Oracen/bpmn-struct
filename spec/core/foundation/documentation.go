package foundation

import (
	"fmt"

	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/validation"
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
		name = fmt.Sprintf("Import:%s", shared.TruncateStringField(d.Text))
	}

	checksBase := d.BaseElement.Validate(name)
	checksCurrent := []error{
		validation.ValNonzero(name, "ImportType", d.Text),
		validation.ValNonzero(name, "Namespace", d.TextFormat),
	}
	return validation.FilterErrors(append(checksBase, checksCurrent...))
}
