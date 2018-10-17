package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Printf("The value is %v\n", m["Answer"])

	delete(m, "Answer")
	fmt.Printf("The value is %v\n", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value ", v, " Present? ", ok)
}
