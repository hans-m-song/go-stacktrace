package errors

import "runtime"

// max callstack depth to return
var PcSize = 10

type frame struct {
	Function string `json:"function"`
	Line     int    `json:"line"`
	File     string `json:"file"`
}

func trace() []frame {
	frames := []frame{}
	pc := make([]uintptr, PcSize)
	skip := 2 // 1 for trace(), 1 for err.Trace()
	count := runtime.Callers(skip, pc)
	stack := runtime.CallersFrames(pc[:count])

	for {
		entry, next := stack.Next()

		if entry.Function != "runtime.goexit" {
			frames = append(frames, frame{
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
