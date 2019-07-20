package decoupling

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// MethodsExample1 is a sample program to show how to declare methods and how the Go
// compiler supports them.
func MethodsExample1() {
	// Values of type user can be used to call methods
	// declared with both value and pointer receivers.
	bill := user{"Bill", "bill@email.com"}
	bill.changeEmail("bill@hotmail.com")
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with both value and pointer receiver.
	joan := &user{"Joan", "joan@email.com"}
	joan.changeEmail("joan@hotmail.com")
	joan.notify()

	// Create a slice of user values with two users.
	users := []user{
		{"ed", "ed@email.com"},
		{"erick", "erick@email.com"},
	}

	// Iterate over the slice of users switching
	// semantics. Not Good!
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}

	// Exception example: Using pointer semantics
	// for a collectoin of strings.
	keys := make([]string, 10)
	for i := range keys {
		keys[i] = func() string { return "key-gen" }()
	}
}

// *****************************************************************************

// duration is a named type that represents a duration
// of time in Nanosecond.
type duration int64

const (
	nanosecond  duration = 1
	microsecond          = 1000 * nanosecond
	millisecond          = 1000 * microsecond
	second               = 1000 * millisecond
	minute               = 60 * second
	hour                 = 60 * minute
)

// setHours sets the specified number of hours.
func (d *duration) setHours(h float64) {
	*d = duration(h) * hour
}

// hours returns the duration as a floating point number of hours.
func (d duration) hours() float64 {
	hour := d / hour
	nsec := d % hour
	return float64(hour) + float64(nsec)*(1e-9/60/60)
}

// MethodsExample2 is a sample program to show how to declare methods against
// a named type.
func MethodsExample2() {
	// Declare a variable of type duration set to
	// its zero value.
	var dur duration

	// Change the value of dur to equal
	// five hours.
	dur.setHours(5)

	// Display the new value of dur.
	fmt.Println("Hours:", dur.hours())
}

// *****************************************************************************

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// setAge sets the age and displays the value.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

// MethodsExample3 is a sample program to show how to declare function variables.
func MethodsExample3() {
	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	fmt.Println("Proper Calls to Methods:")

	// How we actually call methods in Go.
	d.displayName()
	d.setAge(45)

	fmt.Println("\nWhat the Compiler is Doing:")

	// This is what Go is doing underneath.
	data.displayName(d)
	(*data).setAge(&d, 45)

	// =========================================================================

	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method
	// is using a value receiver.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Change the value of d.
	d.name = "Joan"

	// Call the method via the variable. We don't see the change.
	f1()

	// =========================================================================

	fmt.Println("\nCall Pointer Receiver Method with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method
	// is using a pointer receiver.
	f2 := d.setAge

	// Call the method via the variable.
	f2(45)

	// Change the value of d.
	d.name = "Sammy"

	// Call the method via the variable. We see the change.
	f2(45)
}

// *****************************************************************************

// event displays global events.
func event(message string) {
	fmt.Println(message)
}

// event displays events for this data.
func (d *data) event(message string) {
	fmt.Println(d.name, message)
}

// fireEvent1 uses an anonymous function type.
func fireEvent1(f func(string)) {
	f("anonymous")
}

// handler represents a function for handling events.
type handler func(string)

// fireEvent2 uses a function type.
func fireEvent2(h handler) {
	h("handler")
}

// MethodsExample4 is a sample program to show how to declare and use function types.
func MethodsExample4() {
	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	// Use the fireEvent1 handler that accepts any
	// function or method with the right signature.
	fireEvent1(event)
	fireEvent1(d.event)

	// Use the fireEvent2 handler that accepts any
	// function or method of type `handler` or any
	// literal function or method with the right signature.
	fireEvent2(event)
	fireEvent2(d.event)

	// Declare a variable of type handler for the
	// global and method based event functions.
	h1 := handler(event)
	h2 := handler(d.event)

	// User the fireEvent2 handler that accepts
	// values of type handler.
	fireEvent2(h1)
	fireEvent2(h2)

	// User the fireEvent1 handler that accepts
	// any function or method with the right signature.
	fireEvent1(h1)
	fireEvent1(h2)
}

// *****************************************************************************

// func ( /* receiver */ ) average() /* return type */ {
// }

// MethodsExercise1 requires to eclare a method that calculates the batting average for a player.
func MethodsExercise1() {
	// Create a slice of players and populate each player
	// with field values.

	// Display the batting average for each player in the slice.
}
