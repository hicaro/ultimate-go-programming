package decoupling

import (
	"fmt"
	"ultimate-go-programming/language/decoupling/packages/counters"
	"ultimate-go-programming/language/decoupling/packages/toy"
	"ultimate-go-programming/language/decoupling/packages/users"
)

// ExportingExample1 is a sample program to show how to access an exported identifier.
func ExportingExample1() {
	// Create a variable of the exported type and initialize the value to 10.
	counter := counters.AlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)

}

// ExportingExample2 is a sample program to show how the program can't access an
// unexported identifier from another package.
func ExportingExample2() {
	// Create a variable of the unexported type and initialize the value to 10.
	// counter := counters.alertCounter(10)

	// ./example2.go:17: cannot refer to unexported name counters.alertCounter
	// ./example2.go:17: undefined: counters.alertCounter

	// fmt.Printf("Counter: %d\n", counter)
}

// ExportingExample3 is a sample program to show how the program can access a value
// of an unexported identifier from another package.
func ExportingExample3() {
	// Create a variable of the unexported type using the exported
	// New function from the package counters.
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}

// ExportingExample4 is a sample program to show how unexported fields from an exported struct
// type can't be accessed directly.
func ExportingExample4() {
	// Create a value of type User from the users package.
	u := users.User{
		Name: "Chole",
		ID:   10,

		// password: "xxxx",
	}

	// ./example4.go:21: unknown users.User field 'password' in struct literal

	fmt.Printf("User: %#v\n", u)
}

// ExportingExample5 is a sample program to show how to create values from exported types with
// embedded unexported types.
func ExportingExample5() {
	// Create a value of type Manager from the users package.
	u := users.Manager{
		Title: "Dev Manager",
	}

	// Set the exported fields from the unexported user inner type.
	u.Name = "Chole"
	u.ID = 10

	fmt.Printf("User: %#v\n", u)
}

// ExportingExercise1 is to reate a package named toy with a single exported struct type named Toy. Add
// the exported fields Name and Weight. Then add two unexported fields named
// onHand and sold. Declare a factory function called New to create values of
// type toy and accept parameters for the exported fields. Then declare methods
// that return and update values for the unexported fields.
//
// Create a program that imports the toy package. Use the New function to create a
// value of type toy. Then use the methods to set the counts and display the
// field values of that toy value.
func ExportingExercise1() {
	// Use the New function from the toy package to create a value of
	// type toy.
	t := toy.New("Monster Truck", 11)

	// Use the methods from the toy value to set some initialize
	// values.
	fmt.Println("On Hand", t.OnHand())
	fmt.Println("On Hand", t.UpdateOnHand(12))
	fmt.Println("On Hand", t.OnHand())

	fmt.Println("Sold", t.Sold())
	fmt.Println("Sold", t.UpdateSold(19))
	fmt.Println("Sold", t.Sold())

	// Display each field separately from the toy value.
	fmt.Println(t.Name, t.Weight)
}
