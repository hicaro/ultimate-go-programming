package syntax

import (
	"fmt"
	"time"
)

// ConstantsExample1 is a sample program to show how to declare constants and their
// implementation in Go.
func ConstantsExample1() {
	// Constants live within the compiler.
	// They have a parallel type system.
	// Compiler can perform implicit conversions of untyped constants.

	// Untyped Constants.
	const ui = 12345    // kind: integer
	const uf = 3.141592 // kind: floating-point

	// Typed Constants still use the constant type system but their precision
	// is restricted.
	const ti int = 12345        // type: int
	const tf float64 = 3.141592 // type: float64

	// ./constants.go:XX: constant 1000 overflows uint8
	// const myUint8 uint8 = 1000

	// Constant arithmetic supports different kinds.
	// Kind Promotion is used to determine kind in these scenarios.

	// Variable answer will of type float64.
	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	fmt.Println(answer)

	// Constant third will be of kind floating point.
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

	// Constant zero will be of kind integer.
	const zero = 1 / 3 // KindInt(1) / KindInt(3)

	// This is an example of constant arithmetic between typed and
	// untyped constants. Must have like types to perform math.
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)
}

// ConstantsExample2 is a sample program to show how constants do have a parallel type system.
func ConstantsExample2() {
	const (
		// Max integer value on 64 bit architecture.
		maxInt = 9223372036854775807

		// Much larger value than int64.
		bigger = 9223372036854775808543522345

		// Will NOT compile
		// biggerInt int64 = 9223372036854775808543522345
	)

	fmt.Println("Will Compile")
}

// ConstantsExample3 is a sample program to show how iota works.
func ConstantsExample3() {
	const (
		A1 = iota // 0 : Start at 0
		B1 = iota // 1 : Increment by 1
		C1 = iota // 2 : Increment by 1
	)

	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota // 0 : Start at 0
		B2        // 1 : Increment by 1
		C2        // 2 : Increment by 1
	)

	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1 // 1 : Start at 0 + 1
		B3            // 2 : Increment by 1
		C3            // 3 : Increment by 1
	)

	fmt.Println("3:", A3, B3, C3)

	const (
		Ldate         = 1 << iota //  1 : Shift 1 to the left 0.  0000 0001
		Ltime                     //  2 : Shift 1 to the left 1.  0000 0010
		Lmicroseconds             //  4 : Shift 1 to the left 2.  0000 0100
		Llongfile                 //  8 : Shift 1 to the left 3.  0000 1000
		Lshortfile                // 16 : Shift 1 to the left 4.  0001 0000
		LUTC                      // 32 : Shift 1 to the left 5.  0010 0000
	)

	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}

/*
// A Duration represents the elapsed time between two instants as
// an int64 nanosecond count. The representation limits the largest
// representable duration to approximately 290 years.
type Duration int64

// Common durations. There is no definition for units of Day or larger
// to avoid confusion across daylight savings time zone transitions.
const (
        Nanosecond  Duration = 1
        Microsecond          = 1000 * Nanosecond
        Millisecond          = 1000 * Microsecond
        Second               = 1000 * Millisecond
        Minute               = 60 * Second
        Hour                 = 60 * Minute
)

// Add returns the time t+d.
func (t Time) Add(d Duration) Time
*/

// ConstantsExample4 is a sample program to show how literal, constant and variables work
// within the scope of implicit conversion.
func ConstantsExample4() {
	// Use the time package to get the current date/time.
	now := time.Now()

	// Subtract 5 nanoseconds from now using a literal constant.
	literal := now.Add(-5)

	// Subtract 5 seconds from now using a declared constant.
	const timeout = 5 * time.Second // time.Duration(5) * time.Duration(1000000000)
	constant := now.Add(-timeout)

	// Subtract 5 nanoseconds from now using a variable of type int64.
	// minusFive := int64(-5)
	minusFive := -5 * time.Nanosecond
	variable := now.Add(minusFive)

	// example4.go:50: cannot use minusFive (type int64) as type time.Duration in argument to now.Add

	// Display the values.
	fmt.Printf("Now     : %v\n", now)
	fmt.Printf("Literal : %v\n", literal)
	fmt.Printf("Constant: %v\n", constant)
	fmt.Printf("Variable: %v\n", variable)
}

// ConstantsExercise1 is an exercise to:
// Declare an untyped and typed constant and display their values.
// Multiply two literal constants into a typed variable and display the value.
func ConstantsExercise1() {
	const (
		// Declare a constant named server of kind string and assign a value.
		server = "localhost"

		// Declare a constant named port of type integer and assign a value.
		port int = 8000
	)

	// Display the value of both server and port.
	fmt.Printf("%v:%d\n\n", server, port)

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	const apples = 6.0
	const people = 36
	peoplePerApple := people / apples

	// Display the value of the variable.
	fmt.Println(peoplePerApple)
}
