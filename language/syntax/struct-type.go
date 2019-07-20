package syntax

import "fmt"

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

type user struct {
	firstNane string
	lastName  string
	age       int16
}

// StructTypeExample1 is a sample program to show how to declare and initialize struct types.
func StructTypeExample1() {
	// Declare a variable of type example set to its
	// zero value.
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using
	// a struct literal.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}

// StructTypeExample2 is a sample program to show how to declare and initialize anonymous struct types.
func StructTypeExample2() {
	// Declare a variable of an anonymous type set
	// to its zero value.
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymous type and init
	// using a struct literal.
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the values.
	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}

// StructTypeExample3 is the example #3
func StructTypeExample3() {
	// Declare a variable of an anonymous type and init
	// using a struct literal.
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Create a value of type example.
	var ex example

	// Assign the value of the unnamed struct type
	// to the named struct type value.
	ex = e

	// Display the values.
	fmt.Printf("%+v\n", ex)
	fmt.Printf("%+v\n", e)
	fmt.Println("Flag", e.flag)
	fmt.Println("Counter", e.counter)
	fmt.Println("Pi", e.pi)
}

// StructTypeExercise1 is
func StructTypeExercise1() {
	// Declare variable of type user and init using a struct literal.
	u := user{
		firstNane: "James",
		lastName:  "Bond",
		age:       45,
	}

	// Display the field values.
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)

	// Declare a variable using an anonymous struct.
	ua := struct {
		firstNane string
		lastName  string
		age       int16
	}{
		firstNane: "Peter",
		lastName:  "Parker",
		age:       35,
	}

	// Display the field values.
	fmt.Printf("%v\n", ua)
	fmt.Printf("%+v\n", ua)
}
