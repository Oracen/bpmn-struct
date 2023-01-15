package validation

import (
	"reflect"

	"github.com/Oracen/bpmn-struct/shared"
)

func ValNonzero[T any](structName, fieldName string, value T) (err error) {
	if reflect.ValueOf(&value).Elem().IsZero() {
		err = newValueError(structName, fieldName, value, errLenArrayNotZeroOrOne)
	}
	return
}

func ArrZeroOne[T any](structName, fieldName string, array []T) (err error) {
	arrayLen := len(array)
	if arrayLen > 1 {
		err = newValueError(structName, fieldName, arrayLen, errLenArrayNotZeroOrOne)
	}
	return
}

func ArrOneOrMore[T any](structName, fieldName string, array []T) (err error) {
	arrayLen := len(array)
	if arrayLen == 0 {
		err = newValueError(structName, fieldName, arrayLen, errLenArrayLTOne)
	}
	return
}

func ArraysMaxCount[T any](structName, fieldName string, maxCount int, arrays ...[]T) (err error) {
	count := 0
	for idx := 0; idx < len(arrays); idx++ {
		if len(arrays[idx]) >= 1 {
			count++
		}
	}
	if count > maxCount {
		err = newValueError(structName, fieldName, count, errNumItemsExceedsCount)
	}
	return
}

func ArrCheckItems[T shared.BPMNStruct](name string, inputField []T) (errors []error) {
	errors = []error{}
	for _, item := range inputField {
		errors = append(errors, item.Validate(name)...)
	}
	return
}
