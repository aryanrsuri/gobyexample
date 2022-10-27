package main

import (
	"fmt"
)
func main() {
	var nums []int = []int{2,3,4}
	var sum int = 0

	for _,num := range nums {
		sum += num
	}
	fmt.Println(sum)

	kvs := map[string]string{"A": "apple", "B":"banana"}
	for key, value := range kvs {
		fmt.Printf("%s --> %s \n", key, value)
	}

	for index, roon := range "golang"  {
		fmt.Println(index, roon)
	}
}