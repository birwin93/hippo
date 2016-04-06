package hippo

type Error struct {
	Code    int
	Message string
}

func NewError(code int, message string) *Error {
	return &Error{code, message}
}
