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

type WorkHandler func() (interface{}, error)

type Work struct {
	sync.Mutex
	skip   int
	worker []WorkHandler
	result *sync.Map
	errors *sync.Map
}

func (w *Work) Do(work ...WorkHandler) {
	w.Lock()
	defer w.Unlock()

	w.worker = append(w.worker, work...)
}

// 执行并发任务
func (w *Work) Run(second ...int64) error {
	w.Lock()
	defer w.Unlock()

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

		for i, fn := range w.worker {
			go func(i int, handler WorkHandler) {
				wg.Add(1)
				defer wg.Done()

				st := time.Now()
				res, err := handler()

				w.result.Store(i, res)
				w.errors.Store(i, err)

				if isTrace {
					if err != nil {
						trace = append(trace, color.Error.Text("\t-> run %d work uptime: %s, error: %v", i, time.Since(st), err))
					} else {
						trace = append(trace, color.Question.Text("\t-> run %d work uptime: %s", i, time.Since(st)))
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
			for i := 0; i < len(w.worker); i++ {
				if _, ok := w.result.Load(i); !ok {
					trace = append(trace, color.Warn.Text("\t-> run %d work timeout", i))
				}
			}
		}
	case <-ch:
		// Ok
	}

	if isTrace {
		caller := debug.GetCaller(w.skip)
		log.Trace("[Work] %s:%d, total started %d work, uptime: %s ...", caller.File, caller.Line, len(w.worker), time.Since(now))
		if len(trace) > 0 {
			fmt.Printf("%s\n", strings.Join(trace, "\n"))
		}
	}

	return err
}

// 获取全部任务错误
func (w *Work) Errors() map[int]error {
	var errs = make(map[int]error)
	for i := 0; i < len(w.worker); i++ {
		res, _ := w.errors.Load(i)
		if res != nil {
			errs[i] = res.(error)
		}
	}
	return errs
}

// 获取任务执行结果
func (w *Work) Result() []interface{} {
	var result []interface{}
	for i := 0; i < len(w.worker); i++ {
		res, _ := w.result.Load(i)
		result = append(result, res)
	}
	return result
}

// 获取任务错误
func (w *Work) GetError(i int) error {
	res, _ := w.errors.Load(i)
	if res != nil {
		return res.(error)
	}
	return nil
}

// 获取任务错误
func (w *Work) GetResult(i int) interface{} {
	res, ok := w.errors.Load(i)
	if !ok {
		return nil
	}
	return res
}

// 实例化并发任务
func NewWorks(work ...WorkHandler) *Work {
	w := &Work{
		skip:   3,
		result: new(sync.Map),
		errors: new(sync.Map),
	}
	w.Do(work...)
	return w
}

// 执行并发任务
func RunWorks(works []WorkHandler, timeout ...int64) ([]interface{}, error) {
	w := &Work{
		skip:   4,
		result: new(sync.Map),
		errors: new(sync.Map),
	}
	w.Do(works...)
	if err := w.Run(timeout...); err != nil {
		return nil, err
	}
	return w.Result(), nil
}
