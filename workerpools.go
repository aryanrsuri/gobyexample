package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan int, res chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, " => started", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, " => finished", j)
		res <- j * 2
	}
}

func main() {
	const tasks = 5
	jobs := make(chan int, tasks)
	res := make(chan int, tasks)

	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, res)
	}

	for j := 1; j <= tasks; j++ {
		jobs <- j
	}

	close(jobs)
	for a := 1; a <= tasks; a++ {
		<-res
	}
}
