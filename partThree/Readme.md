## 01 Pointers

Go has pointers. A pointer holds the memory address of a value.

The type *T is a pointer to a T value. Its zero value is nil.

	var p *int
The & operator generates a pointer to its operand.

	i := 42
	p = &i

The * operator denotes the pointer's underlying value.

	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer p

This is known as "dereferencing" or "indirecting".

Unlike C, Go has no pointer arithmetic.


## 02 Structs

A struct is a collection of fields.

	type Vertex struct {
		X int
		Y int
	}

Struct fields are accessed using a dot.

	v := Vertex{1, 2}
	v.X = 4

### Pointers to structs
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

	p := &v
	p.X = 42

### Struct Literals

A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.

## 03 Arrays

The type `[n]T` is an array of `n` values of type `T`.

The expression

	var a [10]int

declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.


## 04 Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

	a[low : high]

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements `1` through `3` of `a`:

	a[1:4]

### Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.



## 05 Slice literals

A slice literal is like an array literal without the length.

This is an array literal:

	[3]bool{true, true, false}

And this creates the same array as above, then builds a slice that references it:

	[]bool{true, true, false}

## 06 Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

	var a [10]int

these slice expressions are equivalent:

	a[0:10]
	a[:10]
	a[0:]
	a[:]


## 07 Slice length and capacity

A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

### Nil slices
The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.


## 08 Creating a slice with make

Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

	a := make([]int, 5)  // len(a)=5

To specify a capacity, pass a third argument to make:

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
