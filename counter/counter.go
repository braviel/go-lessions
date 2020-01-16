package counter

import "sync"

//Counter struct
type Counter struct {
	mu    sync.Mutex
	value int
}

//NewCounter ctor
func NewCounter() *Counter {
	return &Counter{}
}

//Inc increases counter
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

//Value get value of Counter
func (c *Counter) Value() int {
	return c.value
}
