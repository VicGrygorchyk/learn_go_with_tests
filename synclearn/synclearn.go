package synclearn

import (
	"sync"
)

type Counter struct {
	mutex sync.Mutex
	value int
}


func (c *Counter) Inc() {
	c.mutex.Lock()
	c.value ++
	c.mutex.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}