package trace

import "encoding/json"

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

	var message string
	if unknown == nil {
		message = "error not provided"
	} else {
		message = unknown.Error()
	}

	return &Error{
		Name:    "UnnamedError",
		Message: message,
	}
}

// Calls `String()` if implemented by the given error, otherwise calls `Error()`
func String(unknown error) string {
	if err, ok := unknown.(stringable); ok {
		return err.String()
	}

	return unknown.Error()
}

// Attempts to marshal an error, returns default error string if not
func Json(unknown error) string {
	err := Guarantee(unknown)

	if serialised, err := json.Marshal(err); err == nil {
		return string(serialised)
	}

	return unknown.Error()
}
