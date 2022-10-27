package main

import "fmt"

func main1() {
	messages := make(chan string)

	go func() { messages <- "pinged" }()
	// send and receives blocks the call execution so a
	// time or watigroup sync is not required
	//msg := <-messages
	fmt.Println(messages)
}

// channel buffering
func main() {

	messages := make(chan string, 2)
	go func() { messages <- "buffered" }()
	go func() { messages <- "channel" }()
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
