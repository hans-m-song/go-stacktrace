package trace

import (
	"fmt"
)

var (
	// check if Error implements interfaces
	_ SerializableError = (*Error)(nil)
	_ TraceableError    = (*Error)(nil)
)

// Struct representing an error.
type Error struct {
	Name    string            `json:"name"`
	Message string            `json:"message"`
	Meta    map[string]string `json:"meta"`
	Stack   []Frame           `json:"stack"`
}

type SerializableError interface {
	GetStack() []string
	GetMeta() []string
	String() string
	Error() string
}

type TraceableError interface {
	Add(key string, value interface{}) *Error
	Trace(err error) *Error
	Tracef(message string, a ...interface{}) *Error
	Clone() *Error
	DeepClone() *Error
}

// Attaches a named value to the metadata of the error.
//
// Note: values get casted into string with `fmt.Sprintf("%+v", value)`.
func (e *Error) Add(key string, value interface{}) *Error {
	e.Meta[key] = fmt.Sprintf("%+v", value)
	return e
}

// Returns the current status of the stacktrace
func (e *Error) GetStack() []string {
	return serialiseStackFrames(e.Stack)
}

// Returns any metadata attached to the error
func (e *Error) GetMeta() []string {
	return serialiseMetaMap(e.Meta)
}

// Returns superficial information about the error (name and message).
//
// Required to satisfy the error interface.
func (e *Error) Error() string {
	return serialiseNameMessage(e.Name, e.Message)
}

// Takes on the message of the given message and traces the callstack.
//
// Note: should be called when returning an error.
//
// Note: will ignore subsequent calls to `Error.Trace` once a stacktrace exists, call `Error.Clone` to create a fresh instance to guarantee a new trace.
func (e *Error) Trace(err error) *Error {
	// error has already been traced
	if len(e.Stack) > 0 {
		return e
	}

	e.Stack = trace()
	e.Message = err.Error()
	return e
}

// Convenience function that wraps `fmt.Errorf`, otherwise behaves the same as `Error.Trace`.
//
// Should be called when returning an error
func (e *Error) Tracef(message string, a ...interface{}) *Error {
	e.Stack = trace()
	e.Message = fmt.Errorf(message, a...).Error()
	return e
}

// Returns everything about the error (name, message, stacktrace, meta)
func (e *Error) String() string {
	return serialiseError(e)
}

// Creates a new instance of the Error without a stack trace.
func (e *Error) Clone() *Error {
	clone := Error{
		Name:    e.Name,
		Message: e.Message,
		Meta:    make(map[string]string, len(e.Meta)),
	}

	for k, v := range e.Meta {
		clone.Meta[k] = v
	}

	return &clone
}

// Creates a deep-copied clone of the Error including stack trace.
func (e *Error) DeepClone() *Error {
	clone := e.Clone()
	clone.Stack = make([]Frame, len(e.Stack))

	for i, entry := range e.Stack {
		clone.Stack[i] = Frame{
			Function: entry.Function,
			Line:     entry.Line,
			File:     entry.File,
		}
	}

	return clone
}

// Create a new instance of `Error` with a name.
func New(name string) *Error {
	return &Error{
		Name:    name,
		Message: "",
		Meta:    map[string]string{},
		Stack:   []Frame{},
	}
}
