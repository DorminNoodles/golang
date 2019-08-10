package main

import "fmt"

func hello() (string, string, string){

	return "Hello", "Go", "World"
}

func main() {

	// i,j := 0, 10
	// f := func() int { return 7 }
	hello, _, world := hello(); // _ for not take the return value

	fmt.Println(hello, world)
}
