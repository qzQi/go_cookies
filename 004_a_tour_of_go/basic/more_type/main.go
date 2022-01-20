package main

import "fmt"

// Vertex is x y
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Printf("%T\n", v)
	p := &v
	(*p).X = 1e9
	p.X = 1
	fmt.Println(v)
}
