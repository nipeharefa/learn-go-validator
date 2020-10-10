package validator

type Constraint interface {
	IsValid(value interface{}) bool
	// GetErrorMessage() error
}

var ConstraintList = map[string]Constraint{
	"not_blank": &NotBlankConstraint{},
	"uuidv4":    &UUIDVConstriant{},
}
