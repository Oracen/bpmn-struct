package validation

import "reflect"

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
	if arrayLen <= 1 {
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
