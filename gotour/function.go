package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func sub(x int, y int) int {
  return x - y
}

func max(x, y int) int {
  return x
}

func main() {
  fmt.Println(add(42, 13))
  fmt.Println(sub(10, 5))
  fmt.Println(max(15, 20))
}

