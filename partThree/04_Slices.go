package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	var s []int = primes[1:4]
	fmt.Println(s)

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)

	names[0] = "Kevin"
	fmt.Println(names)
}
