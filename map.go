package main

import "fmt"

func main() {
	m := make(map[string] int)

	m["1"] = 1
	m["2"] = 2

	fmt.Println(m)
	delete(m, "2")
	fmt.Println(m)


	_, present := m["2"]
	fmt.Println(present)
}
