package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Print("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
