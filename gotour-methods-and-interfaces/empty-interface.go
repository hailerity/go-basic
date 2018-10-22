package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 15
	describe(i)

	i = "Hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}
