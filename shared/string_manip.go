package shared

import (
	"github.com/Oracen/bpmn-struct/constants"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TruncateStringField(field string) string {
	return field[:minInt(constants.IdentifierTruncationLen, len(field))]
}
