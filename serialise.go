package errors

import (
	"fmt"
	"strings"
)

func serialiseMetaMap(meta map[string]string) []string {
	fragments := []string{}
	for key, value := range meta {
		fragment := fmt.Sprintf("%s=\"%s\"", key, value)
		fragments = append(fragments, fragment)
	}

	return fragments
}

func serialiseStackFrames(stack []Frame) []string {
	frames := []string{}
	for _, frame := range stack {
		line := fmt.Sprintf("%v @ %v:%v", frame.Function, frame.File, frame.Line)
		frames = append(frames, line)
	}

	return frames
}

func serialiseNameMessage(name, message string) string {
	return fmt.Sprintf("%s: %s", name, message)
}

func serialiseError(err *Error) string {
	result := fmt.Sprintf("Error: %s", err.Error())

	stack := err.GetStack()
	if len(stack) > 1 {
		result += "\nStacktrace:\n\t" + strings.Join(stack, "\n\t")
	}

	meta := err.GetMeta()
	if len(meta) > 1 {
		result += "\nMetadata:\n\t" + strings.Join(meta, "\n\t")
	}

	return result
}
