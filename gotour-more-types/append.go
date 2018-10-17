package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// the slice grow as needed.
	s = append(s, 1)
	printSlice(s)

	// we can add more than one elements at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("cap %d len %d %v\n", cap(s), len(s), s)
}
