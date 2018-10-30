package main

import "golang.org/x/tour/tree"
import "fmt"

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch := make(chan int)
	go Walk(t1, ch)
	go Walk(t2, ch)

	return true
}

func main() {
	t1 := tree.New(1)
	ch := make(chan int)
	//t2 := tree.New(2)
	//fmt.Println(t1.String())
	// fmt.Println(Same(tree.New(1), tree.New(2)))
	go Walk(t1, ch)
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
