package errors

import "runtime"

// Max callstack depth to return
var PcSize = 10

// Struct containing information about a stackframe
type Frame struct {
	Function string `json:"function"`
	Line     int    `json:"line"`
	File     string `json:"file"`
}

func trace() []Frame {
	frames := []Frame{}
	pc := make([]uintptr, PcSize)
	skip := 2 // 1 for trace(), 1 for err.Trace()
	count := runtime.Callers(skip, pc)
	stack := runtime.CallersFrames(pc[:count])

	for {
		entry, next := stack.Next()

		if entry.Function != "runtime.goexit" {
			frames = append(frames, Frame{
				Function: entry.Function,
				Line:     entry.Line,
				File:     entry.File,
			})
		}

		if !next || len(frames) >= PcSize {
			break
		}
	}

	return frames
}
