package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.1232, -78.2345,
	},
	"Google": Vertex{
		37.1234, -122.3545,
	},
}

func main() {
	fmt.Println(m)
}
