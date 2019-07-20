package syntax

import (
	"fmt"
	"math"
)

// VariableExample1 is a sample program to show how to declare variables.
func VariableExample1() {
	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Specify type and perform a conversion.
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}

// VariableExercise1 #1
func VariableExercise1() {
	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c bool

	// Display the value of those variables.
	fmt.Printf("var a \t %T [%v]\n", a, a)
	fmt.Printf("var b \t %T [%v]\n", b, b)
	fmt.Printf("var c \t %T [%v]\n\n", c, c)

	// Declare variables and initialize.
	// Using the short variable declaration operator.

	aa := 21
	bb := "this is it"
	cc := false

	// Display the value of those variables.
	fmt.Printf("aa \t %T [%v]\n", aa, aa)
	fmt.Printf("bb \t %T [%v]\n", bb, bb)
	fmt.Printf("cc \t %T [%v]\n\n", cc, cc)

	// Perform a type conversion.
	ddd := float32(math.Pi)

	// Display the value of that variable.
	fmt.Printf("ddd \t %T [%v]\n", ddd, ddd)
}
