package validator

type UUIDVConstriant struct {
}

func (uu *UUIDVConstriant) IsValid(value interface{}) bool {
	return false
}
