package main

import "fmt"

type rectangle struct {
	width  int
	height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	R := rectangle{
		width:  10,
		height: 5,
	}

	fmt.Println(R.area(), R.perimeter())

}
