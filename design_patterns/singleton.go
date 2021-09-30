// 单例模式: 一个类只有一个实例且该实例易于外界访问
package design_patterns

import (
	"sync"
)

var (
	instance *Counter
	once     sync.Once
)

type Counter struct {
	number int
	mux    sync.RWMutex
}

func (c *Counter) Add(n int) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.number += n
}

func (c *Counter) Sub(n int) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.number -= n
}

func (c *Counter) Get() int {
	c.mux.RLock()
	defer c.mux.RUnlock()

	return c.number
}

func GetInstance() *Counter {
	once.Do(func() {
		instance = &Counter{
			number: 0,
		}
	})

	return instance
}
