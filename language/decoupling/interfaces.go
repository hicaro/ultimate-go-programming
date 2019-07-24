package decoupling

import (
	"fmt"
	"log"
	"unsafe"
)

// file defines a system file.
type file struct {
	name string
}

// read implements the reader interface for a file.
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

// pipe defines a named pipe network connection.
type pipe struct {
	name string
}

// read implements the reader interface for a network connection.
func (pipe) read(b []byte) (int, error) {
	s := `{name: "bill", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

// retrieveFile can read from a file and process the data.
func retrieveFile(f file) error {
	data := make([]byte, 100)

	len, err := f.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

// retrievePipe can read from a pipe and process the data.
func retrievePipe(p pipe) error {
	data := make([]byte, 100)

	len, err := p.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

// InterfacesExample0 is a sample program that could benefit from polymorphic behavior with interfaces.
func InterfacesExample0() {
	// Create two values one of type file and one of type pipe.
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call each retrieve function for each concrete type.
	retrieveFile(f)
	retrievePipe(p)
}

// *****************************************************************************

// reader is an interface that defines the act of reading data.
type reader interface {
	read(b []byte) (int, error)
}

// retrieve can read any device and process the data.
func retrieve(r reader) error {
	data := make([]byte, 100)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

// InterfacesExample1 is a sample program to show how polymorphic behavior with interfaces.
func InterfacesExample1() {
	// Create two values one of type file and one of type pipe.
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call the retrieve function for each concrete type.
	retrieve(f)
	retrieve(p)
}

// *****************************************************************************

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// person defines a person in the program.
type person struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (p *person) notify() {
	fmt.Printf("Sending Person Email To %s<%s>\n",
		p.name,
		p.email)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}

// InterfacesExample2 is a sample program to show how to understand method sets.
func InterfacesExample2() {
	// Create a value of type person and send a notification.
	p := person{"Bill", "bill@email.com"}

	// Values of type person do not implement the interface because pointer
	// receivers don't belong to the method set of a value.

	sendNotification(&p)

	// sendNotification(p)

	// ./example1.go:36: cannot use p (type person) as type notifier in argument to sendNotification:
	//  person does not implement notifier (notify method has pointer receiver)
}

// *****************************************************************************

// duration is a named type with a base type of int.
// type duration int - Already exists in methods.go

// notify implements the notifier interface.
func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

// InterfacesExample3 is a sample program to show how you can't always get the address of a value.
func InterfacesExample3() {
	// duration(42).notify()

	// 42 is a constant, which is stored in the stack (as opposed to the heap).
	// Thus, it has not addresss to be shared

	// ./example3.go:18: cannot call pointer method on duration(42)
	// ./example3.go:18: cannot take the address of duration(42)
}

// *****************************************************************************

// printer displays information.
type printer interface {
	print()
}

// employee defines a employee in the program.
type employee struct {
	id   int
	name string
}

// print displays the employee's name.
func (e employee) print() {
	fmt.Printf("Employee Name: %s\n", e.name)
}

// InterfacesExample4 is a sample program to show how the concrete value assigned to
// the interface is what is stored inside the interface.
func InterfacesExample4() {
	// Create values of type employee and admin.
	e := employee{12, "Bill"}

	// Add the values and pointers to the slice of
	// printer interface values.
	entities := []printer{
		// Store a copy of the employee value in the interface value.
		e,
		// Store a copy of the address of the employee value in the interface value.
		&e,
	}

	// Change the name field on the employee value.
	e.name = "Bill_CHG"

	// Iterate over the slice of entities and call
	// print against the copied interface value.
	for _, en := range entities {
		en.print()
	}

	// When we store a value, the interface value has its own
	// copy of the value. Changes to the original value will
	// not be seen.

	// When we store a pointer, the interface value has its own
	// copy of the address. Changes to the original value will
	// be seen.
}

// *****************************************************************************

// finder represents the ability to find employees.
type finder interface {
	find(id int) (*employee, error)
}

// employeeSVC is a service for dealing with employees.
type employeeSVC struct {
	host string
}

// find implements the finder interface using pointer semantics.
func (*employeeSVC) find(id int) (*employee, error) {
	return &employee{id: id, name: "Anna Walker"}, nil
}

// InterfacesExample5 is a sample program to show the syntax of type assertions.
func InterfacesExample5() {
	// run performs the find operation against the concrete data that
	// is passed into the call.
	run := func(f finder) error {
		u, err := f.find(1234)
		if err != nil {
			return err
		}
		fmt.Printf("Found employee %+v\n", u)

		// Ideally the finder abstraction would encompass all of
		// the behavior you care about. But what if, for some reason,
		// you really need to get to the concrete value stored inside
		// the interface?

		// Can you access the "host" field from the concrete employeeSVC type pointer
		// that is stored inside this interface variable? No, not directly.
		// All you know is the data has a method named "find".

		// ./example5.go:61:26: f.host undefined (type finder has no field or method host)
		// log.Println("queried", f.host)

		// You can use a type assertion to get a copy of the employeeSVC pointer
		// that is stored inside the interface.
		svc := f.(*employeeSVC)
		log.Println("queried", svc.host)

		return nil
	}

	svc := employeeSVC{
		host: "localhost:3434",
	}

	if err := run(&svc); err != nil {
		log.Fatal(err)
	}
}

// *****************************************************************************

// mockSVC defines a mock service we will access.
type mockSVC struct{}

// find implements the finder interface using pointer semantics.
func (*mockSVC) find(id int) (*employee, error) {
	return &employee{id: id, name: "Jacob Walker"}, nil
}

// InterfacesExample6 is a sample program to show type assertions using the comma-ok idiom.
func InterfacesExample6() {
	run := func(f finder) error {
		u, err := f.find(1234)
		if err != nil {
			return err
		}
		fmt.Printf("Found employee %+v\n", u)

		// If the concrete type value stored inside the interface value is of the
		// type *employeeSVC, then "ok" will be true and "svc" will be a copy of the
		// pointer stored inside the interface.
		if svc, ok := f.(*employeeSVC); ok {
			log.Println("queried", svc.host)
		}

		return nil
	}
	var svc mockSVC

	if err := run(&svc); err != nil {
		log.Fatal(err)
	}
}

// *****************************************************************************

// InterfacesExample7 is a sample program to show the syntax and mechanics of type
// switches and the empty interface.
func InterfacesExample7() {
	myPrintln := func(a interface{}) {
		switch v := a.(type) {
		case string:
			fmt.Printf("Is string  : type(%T) : value(%s)\n", v, v)
		case int:
			fmt.Printf("Is int     : type(%T) : value(%d)\n", v, v)
		case float64:
			fmt.Printf("Is float64 : type(%T) : value(%f)\n", v, v)
		default:
			fmt.Printf("Is unknown : type(%T) : value(%v)\n", v, v)
		}
	}

	// fmt.Println can be called with values of any type.
	fmt.Println("Hello, world")
	fmt.Println(12345)
	fmt.Println(3.14159)
	fmt.Println(true)

	// How can we do the same?
	myPrintln("Hello, world")
	myPrintln(12345)
	myPrintln(3.14159)
	myPrintln(true)

	// - An interface is satisfied by any piece of data when the data exhibits
	// the full method set of behavior defined by the interface.
	// - The empty interface defines no method set of behavior and therefore
	// requires no method by the data being stored.

	// - The empty interface says nothing about the data stored inside
	// the interface.
	// - Checks would need to be performed at runtime to know anything about
	// the data stored in the empty interface.
	// - Decouple around well defined behavior and only use the empty
	// interface as an exception when it is reasonable and practical to do so.
}

// *****************************************************************************

// notify implements the notifier interface.
func (e employee) notify() {
	fmt.Println("Alert", e.name)
}

// InterfacesAdvancedExample1 is a sample program that explores how interface assignments work when
// values are stored inside the interface.
func InterfacesAdvancedExample1() {

	inspect := func(n *notifier, u *employee) {
		word := uintptr(unsafe.Pointer(n)) + uintptr(unsafe.Sizeof(&u))
		value := (**employee)(unsafe.Pointer(word))
		fmt.Printf("Addr employee: %p  Word Value: %p  Ptr Value: %v\n", u, *value, **value)
	}

	// Create a notifier interface and concrete type value.
	var n1 notifier
	u := employee{321, "bill"}

	// Store a copy of the employee value inside the notifier
	// interface value.
	n1 = u

	// We see the interface has its own copy.
	// Addr employee: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
	inspect(&n1, &u)

	// Make a copy of the interface value.
	n2 := n1

	// We see the interface is sharing the same value stored in
	// the n1 interface value.
	// Addr employee: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
	inspect(&n2, &u)

	// Store a copy of the employee address value inside the
	// notifier interface value.
	n1 = &u

	// We see the interface is sharing the u variables value
	// directly. There is no copy.
	// Addr employee: 0x1040a120  Word Value: 0x1040a120  Ptr Value: {bill}
	inspect(&n1, &u)
}

// *****************************************************************************

// Declare the speaker interface with a single method called speak.
type speaker interface {
	speak() string
}

// Declare an empty struct type named english.
type english struct{}

// Declare a method named speak for the english type
// using a value receiver. "Hello World"
func (e english) speak() string {
	return "Hello World"
}

// Declare an empty struct type named chinese.
type chinese struct{}

// Declare a method named speak for the chinese type
// using a pointer receiver. "你好世界"
func (c chinese) speak() string {
	return "你好世界"
}

// sayHello accepts values of the speaker type.
func sayHello(s speaker) {
	// Call the speak method from the speaker parameter.
	fmt.Println(s.speak())
}

// InterfacesExercise1 is supposed to declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
func InterfacesExercise1() {
	// Declare a variable of the interface speaker type
	// set to its zero value.
	var s speaker

	// Declare a variable of type english.
	var e english

	// Assign the english value to the speaker variable.
	s = e

	// Call the speak method against the speaker variable.
	fmt.Println(s.speak())

	// Declare a variable of type chinese.
	var c chinese

	// Assign the chinese pointer to the speaker variable.
	s = c

	// Call the speak method against the speaker variable.
	fmt.Println(c.speak())

	fmt.Println()

	// Call the sayHello function with new values and pointers
	// of english and chinese.
	sayHello(e)
	sayHello(&e)

	fmt.Println()

	sayHello(c)
	sayHello(&c)
}
