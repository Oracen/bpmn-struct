package foundation

import (
	"bpmn-struct/constants"
	"bpmn-struct/validation"
	"fmt"
)

type Documentation struct {
	Text       string `xml:"text" json:"text"`
	TextFormat string `xml:"textFormat" json:"textFormat"`
}

func CreateDocumentation(text string) Documentation {
	return Documentation{
		Text:       text,
		TextFormat: "text/plain",
	}
}

func (d Documentation) Validate() (errors []error) {
	name := fmt.Sprintf("Import:%s", d.Text[:constants.IdentifierTruncationLen])
	checks := []error{
		validation.ValNonzero(name, "ImportType", d.Text),
		validation.ValNonzero(name, "Namespace", d.TextFormat),
	}
	return validation.FilterErrors(checks)
}
