# Go Part Two

## 01 For
For
Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

**Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.**

The init and post statements are optional.

	sum := 1
	for ; sum < 1000 {
		sum += sum
	}

At that point you can drop the semicolons: **C's while is spelled for in Go.**

	for sum < 1000 {
		sum += sum
	}

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

	for {

	}

## 02 If

Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

### If with a short statement

Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

	if (v := math.Pow(x, n); v < 42) {
		v += 82 //v exist only in IF scope
	}

### If and else

Variables declared inside an if short statement are also available inside any of the else blocks.


## 03 Switch

A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.


	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("%s\n", os)
	}

### Switch with no condition

Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains.

	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
	}


## 04 Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	func main() {
		defer fmt.Println("world")

		fmt.Println("hello")
	}
	//hello world

### Stacking defers

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	//9
	//8
	//7
	//6
	//5
	//4
	//3
	//2
	//1
	//0
