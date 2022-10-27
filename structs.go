package main

import "fmt"

type person struct {
	name string
	age  int
}

func person_init(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	s := person{
		name: "Aryan",
		age:  20,
	}

	S := person_init("ARYAN")
	fmt.Println(s, &s, S, &S, *S)
}
