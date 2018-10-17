package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.1233, -78.2343},
	"Google":    {23.8023, -122.3121},
}

func main() {
	fmt.Println(m)
}
