package main

import (
	"fmt"
	"time"
)

// channels can be used to synch execution using a blocking <-channel
func worker(done chan bool) {

	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	d := make(chan bool)
	go worker(d)

	<-d
}
