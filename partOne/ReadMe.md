# Golang Start

## 01 Hello World

Say Hello Go !

## 02 Import Packages

You can import like this :

	import "fmt"
	import "math"

But prefer import like this :

	import (
		"fmt"
		"math"
	)

## 03 Exported Names

In Go, a name is exported if it begins with a capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the math package.

`pizza` and `pi` do not start with a capital letter, so they are not exported.
When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

	fmt.Println(math.Pi) //Pi is exported
	fmt.Println(math.pi) //but not pi


## 04 Functions

A function can take zero or more arguments
In this example, add takes two parameters of type int.
Notice that the type comes after the variable name.

	func add(x int, y int) int {
		return x + y
	}

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In this example, we shortened

	x int, y int
to

	x, y int
