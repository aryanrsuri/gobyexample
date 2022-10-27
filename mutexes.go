package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) incr(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	var wg sync.WaitGroup
	container := Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	increment := func(name string, n int) {
		for i := 0; i < n; i++ {
			container.incr(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go increment("a", 10000)
	go increment("b", 10000)
	go increment("a", 10000)

	wg.Wait()
	fmt.Println(container.counters)
}
