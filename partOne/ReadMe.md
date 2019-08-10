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

## 09 Short Variable Declarations

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct **is not available**.

	var i, j, k int = 1, 2, 3
	i, j, k := 1, 2, 3

## 10 Basic Types

Go's Basic Types are

	bool

	string

	int	int8	int16	int32	int64
	uint	uint8	uint16	uint32 uint64 uintptr

	byte //alias for uint8

	rune // alias for int32
		// a Unicode code point

	float32 float64

	complex64 complex128

## 11 Zero Value

Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.


## 12 Type conversions

The expression `T(v)` converts the value `v` to the type `T`.
Some numeric conversions:

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

Or, put more simply:

	i := 42
	f := float64(i)
	u := uint(f)

**Unlike in C, in Go assignment between items of different type requires an explicit conversion.**


## 13 Type Inference

When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

	var i int
	j := i // j is an int
But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

## 14 Constants

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

	cont World = "世界"

**Constants cannot be declared using the := syntax.**
