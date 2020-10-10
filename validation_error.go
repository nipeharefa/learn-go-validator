package validator

type ErrorField struct {
	FieldName string `json:"fieldName"`
}

type ErrorBag struct {
	errorFields []ErrorField
}

func NewErrorBag() *ErrorBag {
	return &ErrorBag{}
}

func (eb *ErrorBag) AddError(err ErrorField) {
	eb.errorFields = append(eb.errorFields, err)
}

func (eb *ErrorBag) HasError() bool {
	return len(eb.errorFields) > 0
}
