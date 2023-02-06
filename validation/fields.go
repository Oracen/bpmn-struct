package validation

import (
	"reflect"

	"github.com/Oracen/bpmn-struct/shared"
	"golang.org/x/exp/constraints"
)

func ValGTEZero[T constraints.Float | constraints.Integer](structName, fieldName string, value T) (err error) {
	if value < 0 {
		err = newValueError(structName, fieldName, value, errValueLTZero)
	}
	return
}

func ArrGTEZero[T constraints.Float | constraints.Integer](structName, fieldName string, array []T) (err []error) {
	err = []error{}
	for _, item := range array {
		err = append(err, ValGTEZero(structName, fieldName, item))
	}
	return
}

func ValNonzero[T any](structName, fieldName string, value T) (err error) {
	if reflect.ValueOf(&value).Elem().IsZero() {
		err = newValueError(structName, fieldName, value, errValueZeroEquivalent)
	}
	return
}

func ArrNonzero[T any](structName, fieldName string, array []T) (err []error) {
	err = []error{}
	for _, item := range array {
		err = append(err, ValNonzero(structName, fieldName, item))
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
