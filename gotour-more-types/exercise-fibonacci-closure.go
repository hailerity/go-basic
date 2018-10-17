package main

import "fmt"

func fibonacci() func() int {
	first := 0
	second := 0

	return func() int {
		if first == 0 {
			first = 1
			return 1
		}

		if second == 0 {
			second = 1
			return 1
		}

		next := first + second
		first = second
		second = next

		return next
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
