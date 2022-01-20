package errors

import (
	"fmt"
)

type Error struct {
	Name    string            `json:"name"`
	Message string            `json:"message"`
	Meta    map[string]string `json:"meta"`
	Stack   []frame           `json:"stack"`
}

// attaches a named value to the metadata of the error
// values get casted into string with `fmt.Sprintf("%+v", value)`
func (e *Error) Add(key string, value interface{}) *Error {
	e.Meta[key] = fmt.Sprintf("%+v", value)
	return e
}

// returns the current status of the stacktrace
func (e *Error) GetStack() []string {
	return serialiseStackFrames(e.Stack)
}

// returns any metadata attached to the error
func (e *Error) GetMeta() []string {
	return serialiseMetaMap(e.Meta)
}

// returns superficial information about the error
// (name and message)
// required to satisfy the error interface
func (e *Error) Error() string {
	return serialiseNameMessage(e.Name, e.Message)
}

// takes on the message of the given message and traces the callstack
// should be called when returning an error
func (e *Error) Trace(err error) *Error {
	e.Stack = trace()
	e.Message = err.Error()
	return e
}

// returns everything about the error
// (name, message, stacktrace, meta)
func (e *Error) String() string {
	return serialiseError(e)
}

// creates a deep-copied clone of the Error
func (e *Error) Clone() *Error {
	clone := Error{
		Name:    e.Name,
		Message: e.Message,
		Meta:    make(map[string]string, len(e.Meta)),
		Stack:   make([]frame, len(e.Stack)),
	}

	for key, value := range e.Meta {
		clone.Meta[key] = value
	}

	for i, entry := range e.Stack {
		clone.Stack[i] = frame{
			Function: entry.Function,
			Line:     entry.Line,
			File:     entry.File,
		}
	}

	return &clone
}

// create a new instance of `Error` with a name
func NewError(name string) *Error {
	return &Error{
		Name:    name,
		Message: "",
		Meta:    map[string]string{},
		Stack:   []frame{},
	}
}
