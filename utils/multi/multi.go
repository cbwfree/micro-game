// 多任务
package multi

import (
	"errors"
	"time"
)

var (
	ErrTimeout  = errors.New("run timeout")
	ErrNotFound = errors.New("not found")
)

const (
	DefaultTimeout  = 5 * time.Second
	DefaultLongTime = 200 * time.Millisecond
)
