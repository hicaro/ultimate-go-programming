package language

import "fmt"

// person represents a person in the system.
type person struct {
	name   string
	email  string
	logins int
}

// Number of elements to grow each stack frame.
// Run with 10 and then with 1024
const size = 10

// PointersExample1 - Sample program to show the basic concept of pass by value.
func PointersExample1() {
	// increment declares count as an integer
	increment := func(inc int) {

		// Increment the "value of" inc.
		inc++
		println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
	}

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	increment(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

// PointersExample2 - Sample program to show the basic concept of using a pointer
// to share data.
func PointersExample2() {
	// increment declares count as a pointer variable whose value is
	// always an address and points to values of type int.
	increment := func(inc *int) {

		// Increment the "value of" count that the "pointer points to".
		*inc++

		println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
	}

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	increment(&count)

	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

// PointersExample3 - Sample program to show the basic concept of using a pointer
// to share data.
func PointersExample3() {
	// increment declares logins as a pointer variable whose value is
	// always an address and points to values of type int.
	increment := func(logins *int) {
		*logins++
		fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n\n", &logins, logins, *logins)
	}

	// display declares u as person pointer variable whose value is always an address
	// and points to values of type person.
	display := func(u *person) {
		fmt.Printf("%p\t%+v\n", u, *u)
		fmt.Printf("Name: %q Email: %q Logins: %d\n\n", u.name, u.email, u.logins)
	}

	// Declare and initialize a variable named bill of type person.
	bill := person{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	//** We don't need to include all the fields when specifying field
	// names with a struct literal.

	// Pass the "address of" the bill value.
	display(&bill)

	// Pass the "address of" the logins field from within the bill value.
	increment(&bill.logins)

	// Pass the "address of" the bill value.
	display(&bill)
}

// PointersExample4 - Sample program to teach the mechanics of escape analysis.
func PointersExample4() {
	// createUserV1 creates a person value and passed
	// a copy back to the caller.
	p1 := func() person {
		p := person{
			name:  "Bill",
			email: "bill@ardanlabs.com",
		}

		println("V1", &p)

		return p
	}()

	// createUserV2 creates a person value and shares
	// the value with the caller.
	p2 := func() *person {
		p := person{
			name:  "Bill",
			email: "bill@ardanlabs.com",
		}

		println("V2", &p)

		return &p
	}()

	println("p1", &p1, "p2", p2)
}

// PointersExample5 - Sample program to show how stacks grow/change.
func PointersExample5() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

// Helper functions

// stackCopy recursively runs increasing the size
// of the stack.
func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}

// PointersExercise1 - Declare and initialize a pointer variable of type int that points to the last
// variable you just created. Display the _address of_ , _value of_ and the
// _value that the pointer points to_.
func PointersExercise1() {
	// Declare an integer variable with the value of 20.

	// Display the address of and value of the variable.

	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.

	// Display the address of, value of and the value the pointer
	// points to.
}

// PointersExercise2 - Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
func PointersExercise2() {
	// Create a function that changes the value of one of the user fields.
	// funcName := func( /* add pointer parameter, add value parameter */ ) {

	// 	// Use the pointer to change the value that the
	// 	// pointer points to.
	// }

	// Create a variable of type user and initialize each field.

	// Display the value of the variable.

	// Share the variable with the function you declared above.

	// Display the value of the variable.
}
