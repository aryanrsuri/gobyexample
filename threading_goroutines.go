package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, "=> ", i)
	}

}

func main() {
	// blocked sequential call
	f("direct")

	// concurrent
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// wait for the calls to finish (more approp to use a WaitGroup)
	time.Sleep(3 * time.Millisecond)
	fmt.Println("done")
}
