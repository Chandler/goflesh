package types

type ValueNotSetError struct {
	Key string // key of the not-set value
}

func (e *ValueNotSetError) Error() string {
	return "ValueNotSetError: Missing value for " + e.Key
}

type DatabaseError struct {
	Message string
}

func (e *DatabaseError) Error() string {
	return "DatabaseError: " + e.Message
}
