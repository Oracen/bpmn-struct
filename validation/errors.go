package validation

import (
	"errors"
	"fmt"
)

var (
	errLenArrayNotZeroOrOne = errors.New("array is should have length 0 or 1")
	errLenArrayLTOne        = errors.New("array is should have at least one object")
	errValueLTZero          = errors.New("numeric value should be greater than or equal to 0")
	errValueZeroEquivalent  = errors.New("value should not be null or zero equivalent")
	errNumItemsExceedsCount = errors.New("number of objects in array should not exceed count")
	errInvalidEnumValue     = errors.New("given enum value does not exist")
)

type valueError[T any] struct {
	Struct string
	Field  string
	Value  T
	Err    error
}

func newValueError[T any](structName, fieldName string, value T, err error) valueError[T] {
	return valueError[T]{
		Struct: structName,
		Field:  fieldName,
		Value:  value,
		Err:    err,
	}
}

func (v valueError[T]) Error() string {
	return fmt.Sprintf("value error on %s:%s: %s, got %v", v.Struct, v.Field, v.Err, v.Value)
}

func FilterErrors(checks []error) (errors []error) {
	for _, item := range checks {
		if item != nil {
			errors = append(errors, item)
		}
	}
	return

}

func NewEnumValidationError[T any](structName, field string, validValues []T, got T) (err error) {
	valid := fmt.Sprintf("%v, valid values %v", got, validValues)
	return newValueError(structName, field, valid, errInvalidEnumValue)
}
