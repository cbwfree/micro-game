package tool

import "sync"

// 计数器
type Counter struct {
	sync.Mutex
	count int64
}

// 设置初始值
func (c *Counter) SetStart(count int64) {
	c.Lock()
	defer c.Unlock()

	c.count = count
}

func (c *Counter) Add(num int64) {
	c.Lock()
	defer c.Unlock()

	c.count += num
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()

	c.count++
}

func (c *Counter) Total() int64 {
	return c.count
}

func NewCounter() *Counter {
	return &Counter{}
}
