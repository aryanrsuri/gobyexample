package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int) {

	fmt.Printf("Worker starting %v\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker finished %v\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			Worker(i)
		}()

		wg.Wait()
	}
}
