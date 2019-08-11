package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Vector struct {
	X, Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}

	v.X = 42
	fmt.Println(v.X)

	pointer := &v
	pointer.Y = 64
	fmt.Println(pointer.Y)

	var (
		v1 = Vector{1, 2}  // has type Vector
		v2 = Vector{X: 1}  // Y:0 is implicit
		v3 = Vector{}      // X:0 and Y:0
		p  = &Vector{1, 2} // has type *Vector
	)

	fmt.Println(v1, p, v2, v3)

}
