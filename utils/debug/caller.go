package debug

import (
	"runtime"
)

type CallerInfo struct {
	File string
	Line int
	Func string
}

// 获取调用信息 (单条)
func GetCaller(skip int) *CallerInfo {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip, rpc)
	if n < 1 {
		return &CallerInfo{}
	}
	frame, _ := runtime.CallersFrames(rpc).Next()
	return &CallerInfo{
		File: frame.File,
		Line: frame.Line,
		Func: frame.Function,
	}
}

// 获取调用信息 (多条)
func GetCallers(skip int, limit int) []*CallerInfo {
	rpc := make([]uintptr, limit)
	n := runtime.Callers(skip, rpc)
	if n < 1 {
		return nil
	}
	frames := runtime.CallersFrames(rpc[:n])
	var lines []*CallerInfo
	for {
		frame, more := frames.Next()
		lines = append(lines, &CallerInfo{
			File: frame.File,
			Line: frame.Line,
			Func: frame.Function,
		})
		if !more {
			break
		}
	}
	return lines
}
