package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func check(x int) (ret int) {

	if v := 42; x < v {
		ret = 7
	} else {
		//v exist in else
		fmt.Println("V is smaller than X, V = ", v);
	}
	return ret
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 20),
	)

	fmt.Println(check(46))
}
