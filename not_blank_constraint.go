package validator

import (
	"reflect"
)

type NotBlankConstraint struct {
	err error
}

func (n NotBlankConstraint) IsValid(v interface{}) bool {

	vv := reflect.ValueOf(v)

	if vv.Type().Kind() != reflect.String {
		return false
	}

	str := vv.Interface().(string)

	return str != ""
}

func (n NotBlankConstraint) GetErrorMessage() error {
	return n.err
}
