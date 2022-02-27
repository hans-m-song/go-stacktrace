package trace

type stringable interface {
	String() string
}

// Create a new instance of `Error` from an error
//
// If `unknown` is of type `Error`, it returns a clone
func Guarantee(unknown error) *Error {
	if err, ok := unknown.(*Error); ok {
		return err
	}

	return &Error{
		Name:    "UnnamedError",
		Message: unknown.Error(),
	}
}

// Calls `String()` if implemented by the given error, otherwise calls `Error()`
func String(unknown error) string {
	if err, ok := unknown.(stringable); ok {
		return err.String()
	}

	return unknown.Error()
}
