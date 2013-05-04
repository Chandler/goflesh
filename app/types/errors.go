package types

type ValueNotSetError struct {
	Key string // key of the not-set value
}

func (e *ValueNotSetError) Error() string {
	return "Missing value for: " + e.Key
}
