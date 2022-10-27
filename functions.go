package main

import "fmt"

func plus (A int , B int ) int {
	return A + B
}

func plusPlus(A int, B int, C int) int {
	return A + B + C
}

func vals() (int, int) {
	return 3,7
}

func sum(nums ...int) int {
	var total int = 0
	for _, num := range nums {
		total += num
	}
	return total
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
// Closures -- anonymous functions
// Recursion -- is supported 

func main() {
	fmt.Println(plus(1,2), plusPlus(1,2,3))
	A, B := vals()
	nums := []int{1,2,3,4}
	fmt.Println(A,B)
	fmt.Println(sum(nums...))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib( n - 1 ) + fib( n - 2 )
	}

	fmt.Println(factorial(7))
	fmt.Println(fib(7))
}