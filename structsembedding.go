package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("(num with base) => %v", b.num)
}

func main() {
	co := container{
		base: base{
			num: 12,
		},
		str: "12",
	}

	fmt.Println(co.num, co.str, co.base)
	fmt.Println(co.describe())

	// Embedding structs with methods may be used to bestow interface implementations onto other structs.
	// Here we see that a container now implements the describer interface because it embeds base.

	type describer interface {
		describe() string
	}

	var d describer = co

	fmt.Println(d.describe())

}
