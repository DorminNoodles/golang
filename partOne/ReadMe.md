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

## 05 Functions Multiple Results

A function can return any number of results.

	func swap(x,y string) (string, string) {
		return y, x
	}

We get the return with the `:=` operator

	a,b := swap("hello", "world")

Link to documentation https://golang.org/ref/spec#Short_variable_declarations

## 06 Named Return Values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

	func split(sum int) (x, y int) {
		x = * 4 / 9
		y = sum - x
		return //x, y
	}

## 07 Variables

The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.

	package main

	import "fmt"

	var c, python, java bool

	func main() {
		var i int
		fmt.Println(i, c, python, java)
	}

## 08 Variables with initializers
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

	var i, j int = 1, 2

	func main() {
		var c, python, java = true, false, "no!"
	}
