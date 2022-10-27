package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	response, error := http.Get("https://gobyexample.com")
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Println(response.Status)
}
