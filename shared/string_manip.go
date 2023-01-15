package shared

import (
	"fmt"
	"reflect"

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

func TypeNameString(name, inputStruct any, identifier string) string {
	if name == "" {
		name = fmt.Sprintf("%s:%s", reflect.ValueOf(inputStruct).Type().Name(), identifier)
	} else {
		name = fmt.Sprintf("%s:%s", name, reflect.ValueOf(inputStruct).Type().Name())
	}
	return name.(string)
}
