package concurrency

import "sync"

type Counter struct {
	count      int
	sync.Mutex // direct embedding
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}
