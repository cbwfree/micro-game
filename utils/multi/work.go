package multi

import (
	"bytes"
	"github.com/cbwfree/micro-game/utils/color"
	"github.com/cbwfree/micro-game/utils/debug"
	"github.com/cbwfree/micro-game/utils/log"
	"sync"
	"time"
)

type WorkHandler func() (interface{}, error)

type WorkResult struct {
	Data  interface{}
	Error error
}

type Work struct {
	sync.Mutex
	wg     *sync.WaitGroup
	result *sync.Map
	worker []WorkHandler
	level  int
}

func (w *Work) Do(work ...WorkHandler) {
	w.worker = append(w.worker, work...)
}

// 执行并发任务
func (w *Work) Run(timeout ...time.Duration) error {
	w.Lock()
	defer w.Unlock()

	var trace *bytes.Buffer
	if log.IsTrace() {
		trace = new(bytes.Buffer)
	}

	now := time.Now()
	done := make(chan struct{})
	workNum := len(w.worker)

	w.result = new(sync.Map)
	w.wg = new(sync.WaitGroup)

	w.wg.Add(workNum)
	for i, fn := range w.worker {
		go func(i int, fn WorkHandler) {
			defer w.wg.Done()
			st := time.Now()
			var res = new(WorkResult)
			res.Data, res.Error = fn()
			w.result.Store(i, res)
			if trace != nil {
				diffTime := time.Since(st)
				if res.Error != nil {
					trace.WriteString(color.Error.Text(" -> Run %d work time: %s, Error: %v\n", i, diffTime, res.Error))
				} else if diffTime > DefaultLongTime {
					trace.WriteString(color.Secondary.Text(" -> Run %d work time: %s\n", i, diffTime))
				}
			}
		}(i, fn)
	}

	go func() {
		w.wg.Wait()
		done <- struct{}{}
	}()

	var afterTime time.Duration
	if len(timeout) > 0 {
		afterTime = timeout[0]
	} else {
		afterTime = DefaultTimeout
	}

	var err error
	select {
	case <-done:
		err = nil
	case <-time.After(afterTime):
		if log.IsTrace() {
			for i := 0; i < len(w.worker); i++ {
				_, ok := w.result.Load(i)
				if !ok && trace != nil {
					trace.WriteString(color.Warn.Text(" -> Warning: Run %d work timeout\n", i))
				}
			}
		}
		err = ErrTimeout
	}

	if trace != nil {
		caller := debug.GetCaller(w.level)
		log.Printf("%s\n", color.Primary.Text("[Task] %s:%d, Start %d Task, Total Time: %s ...", caller.File, caller.Line, workNum, time.Since(now)))
		if trace.Len() > 0 {
			log.Printf("%s", trace)
		}
	}

	return err
}

// 仅获取数据结果
func (w *Work) Data() []interface{} {
	var result []interface{}
	for i := 0; i < len(w.worker); i++ {
		if res, ok := w.result.Load(i); ok {
			result = append(result, res.(*WorkResult).Data)
		} else {
			result = append(result, nil)
		}
	}
	return result
}

// 仅获取错误
func (w *Work) Errors() []error {
	var errs []error
	for i := 0; i < len(w.worker); i++ {
		if res, ok := w.result.Load(i); ok {
			errs = append(errs, res.(*WorkResult).Error)
		} else {
			errs = append(errs, ErrNotFound)
		}
	}
	return errs
}

// 获取任务执行结果
func (w *Work) Result() []*WorkResult {
	var result []*WorkResult
	for i := 0; i < len(w.worker); i++ {
		if res, ok := w.result.Load(i); ok {
			result = append(result, res.(*WorkResult))
		} else {
			result = append(result, &WorkResult{
				Error: ErrNotFound,
			})
		}
	}
	return result
}

// 执行并获取数据
func (w *Work) RunAndData(timeout ...time.Duration) ([]interface{}, error) {
	w.level = 4
	if err := w.Run(timeout...); err != nil {
		return nil, err
	}
	return w.Data(), nil
}

// 执行并获取结果
func (w *Work) RunAndResult(timeout ...time.Duration) ([]*WorkResult, error) {
	w.level = 4
	if err := w.Run(timeout...); err != nil {
		return nil, err
	}
	return w.Result(), nil
}

// 实例化并发任务
func NewWorks(work ...WorkHandler) *Work {
	w := &Work{
		level: 3,
	}
	w.Do(work...)
	return w
}

// 执行并发任务
func RunWorks(works []WorkHandler, timeout ...time.Duration) ([]*WorkResult, error) {
	w := &Work{
		level: 4,
	}
	w.Do(works...)
	if err := w.Run(timeout...); err != nil {
		return nil, err
	}
	return w.Result(), nil
}
