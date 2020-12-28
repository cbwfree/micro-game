package multi

import (
	"context"
	"fmt"
	"github.com/cbwfree/micro-game/utils/color"
	"github.com/cbwfree/micro-game/utils/debug"
	"github.com/cbwfree/micro-game/utils/errors"
	"github.com/cbwfree/micro-game/utils/log"
	"strings"
	"sync"
	"time"
)

type TaskHandler func() error

type Task struct {
	sync.Mutex
	skip   int
	tasks  []TaskHandler
	errors *sync.Map
}

func (t *Task) Do(work ...TaskHandler) {
	t.Lock()
	defer t.Unlock()

	t.tasks = append(t.tasks, work...)
}

// 执行并发任务
func (t *Task) Run(second ...int64) error {
	t.Lock()
	defer t.Unlock()

	var timeout time.Duration
	if len(second) > 0 {
		timeout = time.Duration(second[0]) * time.Second
	} else {
		timeout = 5 * time.Second
	}

	var trace []string
	var isTrace = log.IsTrace()

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	now := time.Now()
	ch := make(chan struct{}, 0)

	go func() {
		wg := new(sync.WaitGroup)

		for i, fn := range t.tasks {
			wg.Add(1)
			go func(i int, fn TaskHandler) {
				defer wg.Done()

				st := time.Now()
				err := fn()
				t.errors.Store(i, err)

				if isTrace {
					if err != nil {
						trace = append(trace, color.Error.Text("\t-> run %d task time: %s, error: %+v", i, time.Since(st), err))
					} else {
						trace = append(trace, color.Question.Text("\t-> run %d task time: %s", i, time.Since(st)))
					}
				}
			}(i, fn)
		}

		wg.Wait()
		ch <- struct{}{}
	}()

	var err error
	select {
	case <-ctx.Done():
		err = errors.Timeout()
		if isTrace {
			for i := 0; i < len(t.tasks); i++ {
				if _, ok := t.errors.Load(i); !ok {
					trace = append(trace, color.Warn.Text("\t-> run %d task timeout", i))
				}
			}
		}
	case <-ch:
		// Ok
	}

	if isTrace {
		caller := debug.GetCaller(t.skip)
		log.Trace("[Task] %s:%d, total started %d task, uptime: %s ...", caller.File, caller.Line, len(t.tasks), time.Since(now))
		if len(trace) > 0 {
			fmt.Printf("%s\n", strings.Join(trace, "\n"))
		}
	}

	return err
}

// 获取全部任务错误
func (t *Task) Errors() map[int]error {
	var errs = make(map[int]error)
	for i := 0; i < len(t.tasks); i++ {
		res, _ := t.errors.Load(i)
		if res != nil {
			errs[i] = res.(error)
		}
	}
	return errs
}

// 获取任务错误
func (t *Task) GetError(i int) error {
	res, _ := t.errors.Load(i)
	if res != nil {
		return res.(error)
	}
	return nil
}

// 实例化并发任务
func NewTasks(task ...TaskHandler) *Task {
	t := &Task{
		skip:   3,
		errors: new(sync.Map),
	}
	t.Do(task...)
	return t
}

// 执行并发任务
func RunTasks(tasks []TaskHandler, timeout ...int64) error {
	t := &Task{
		skip:   4,
		errors: new(sync.Map),
	}
	t.Do(tasks...)
	return t.Run(timeout...)
}
