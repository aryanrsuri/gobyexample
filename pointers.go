package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeropointer(iptr *int) {
	*iptr = 0
}

func main() {
	var i int = 1
	fmt.Println(i)
	zeroval(i)
	fmt.Println(i)
	zeropointer(&i)
	fmt.Println(i)
	fmt.Println(&i)
}