// https://tour.golang.org/concurrency/9
package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v    int
	lock sync.Mutex
}

func (c *Counter) Inc() {
	c.lock.Lock()
	c.v++
	c.lock.Unlock()
}

func (c *Counter) Read() {
	c.lock.Lock()
	fmt.Println("v", c.v)
	defer c.lock.Unlock()
}

func main() {
	c := Counter{v: 0}

	for i := 0; i < 100; i++ {
		go c.Inc()
		go c.Read()
	}

	time.Sleep(time.Second)
	c.Read()
}
