package errors

type clonable interface {
	Clone() *Error
}

// create a new instance of `Error` from an error
// if `unknown` is of type `Error`, it returns a clone
func Guarantee(unknown error) *Error {
	if err, ok := unknown.(clonable); ok {
		result := err.Clone()
		return result
	}

	return &Error{
		Name:    "UnnamedError",
		Message: unknown.Error(),
		Meta:    map[string]string{},
		Stack:   []frame{},
	}
}
