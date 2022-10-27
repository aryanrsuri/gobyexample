package main

import (
	"fmt"
	"strconv"
)

// number parsing
func nmain() {
	f, _ := strconv.ParseFloat("1.23", 64)
	g, _ := strconv.ParseInt("3", 0, 64)
	h, _ := strconv.ParseInt("0x1c8", 0, 64)
	k, _ := strconv.Atoi("135")

	fmt.Println(f, g, h, k)
}

func main() {
	nmain()
}
