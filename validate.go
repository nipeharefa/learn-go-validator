package validator

import (
	"reflect"
	"strings"
)

// type fieldCollector

func Validate(p interface{}) *ErrorBag {

	errorBag := NewErrorBag()

	v := reflect.ValueOf(p).Elem()
	typeOfT := v.Type()

	numField := v.NumField()
	for i := 0; i < numField; i++ {
		f := v.Field(i)

		field := typeOfT.Field(i)
		tagString, ok := typeOfT.Field(i).Tag.Lookup("vd")
		if !ok {
			continue
		}

		tagValues := strings.Split(tagString, ",")

		for _, tagValue := range tagValues {
			if constraint, ok := ConstraintList[tagValue]; ok {
				isValid := constraint.IsValid(f.Interface())
				if !isValid {
					errorBag.AddError(ErrorField{
						FieldName: field.Name,
					})
				}
			}
		}
	}

	return errorBag
}
