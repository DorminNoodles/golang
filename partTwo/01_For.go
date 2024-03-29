package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	//optional statements
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	//for is while in Go
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
