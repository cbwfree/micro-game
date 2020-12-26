package multi

import (
	"bytes"
	"github.com/cbwfree/micro-game/utils/color"
	"github.com/cbwfree/micro-game/utils/debug"
	"github.com/cbwfree/micro-game/utils/log"
	"sync"
	"time"
)

type TaskHandler func() error

type Task struct {
	sync.Mutex
	wg     *sync.WaitGroup
	tasks  []TaskHandler
	errors *sync.Map
	level  int
}

func (mt *Task) Do(work ...TaskHandler) {
	mt.tasks = append(mt.tasks, work...)
}

// 执行并发任务
func (mt *Task) Run(timeout ...time.Duration) error {
	mt.Lock()
	defer mt.Unlock()

	var trace *bytes.Buffer
	if log.IsTrace() {
		trace = new(bytes.Buffer)
	}

	now := time.Now()
	done := make(chan struct{})
	taskNum := len(mt.tasks)

	mt.errors = new(sync.Map)
	mt.wg = new(sync.WaitGroup)

	mt.wg.Add(taskNum)

	for i, fn := range mt.tasks {
		go func(i int, fn TaskHandler) {
			defer mt.wg.Done()
			st := time.Now()
			err := fn()
			mt.errors.Store(i, err)
			if trace != nil {
				diffTime := time.Since(st)
				if err != nil {
					trace.WriteString(color.Error.Text(" -> Run %d task time: %s, Error: %v\n", i, diffTime, err))
				} else if diffTime > DefaultLongTime {
					trace.WriteString(color.Secondary.Text(" -> Run %d task time: %s\n", i, diffTime))
				}
			}
		}(i, fn)
	}

	go func() {
		mt.wg.Wait()
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
			for i := 0; i < len(mt.tasks); i++ {
				_, ok := mt.errors.Load(i)
				if !ok && trace != nil {
					trace.WriteString(color.Warn.Text(" -> Warning: Run %d task timeout\n", i))
				}
			}
		}
		err = ErrTimeout
	}

	if trace != nil {
		caller := debug.GetCaller(mt.level)
		log.Printf("%s\n", color.Primary.Text("[Task] %s:%d, Start %d Task, Total Time: %s ...", caller.File, caller.Line, taskNum, time.Since(now)))
		if trace.Len() > 0 {
			log.Printf("%s", trace)
		}
	}

	return err
}

// 获取任务执行结果
func (mt *Task) Errors() []error {
	var errs []error
	for i := 0; i < len(mt.tasks); i++ {
		if err, ok := mt.errors.Load(i); ok {
			if err != nil {
				errs = append(errs, err.(error))
			} else {
				errs = append(errs, nil)
			}
		} else {
			errs = append(errs, ErrNotFound)
		}
	}
	return errs
}

// 实例化并发任务
func NewTasks(task ...TaskHandler) *Task {
	w := &Task{
		level: 3,
	}
	w.Do(task...)
	return w
}

// 执行并发任务
func RunTasks(tasks []TaskHandler, timeout ...time.Duration) error {
	w := &Task{
		level: 4,
	}
	w.Do(tasks...)
	return w.Run(timeout...)
}
