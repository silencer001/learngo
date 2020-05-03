package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.0, -74},
	"Google":    {37.4, -122.0},
}

func main() {
	fmt.Println(m)
}
