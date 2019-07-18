package language

import (
	"fmt"
	"unicode/utf8"
)

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			s)
	}
}

type sliceUser struct {
	id    int
	name  string
	likes int
}

// SlicesExample1 is a sample program to show how the capacity of the slice
// is not available for use.
func SlicesExample1() {
	// Create a slice with a length of 5 elements.
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// You can't access an index of a slice beyond its length.
	fruits[5] = "Runtime error"

	// Error: panic: runtime error: index out of range

	fmt.Println(fruits)
}

// SlicesExample2 is a sample program to show the components of a slice. It has a
// length, capacity and the underlying array.
func SlicesExample2() {
	// Create a slice with a length of 5 elements and a capacity of 8.
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	inspectSlice(fruits)
}

// SlicesExample3 is a sample program to show how to takes slices of slices to create different
// views of and make changes to the underlying array.
func SlicesExample3() {
	// Create a slice with a length of 5 elements and a capacity of 8.
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	inspectSlice(slice1)

	// Take a slice of slice1. We want just indexes 2 and 3. [a:b)
	// Parameters are [starting_index : (starting_index + length)]
	slice2 := slice1[2:4]
	// slice2 := slice1[2:4:4] // set capacity to 2 as well
	inspectSlice(slice2)

	fmt.Println("\n*************************")

	// Change the value of the index 0 of slice2.
	slice2[0] = "CHANGED"
	// append(slice2, "CHANGED") // also modifies the original slice (same underlying array)

	// Display the change across all existing slices.
	inspectSlice(slice1)
	inspectSlice(slice2)

	fmt.Println("\n*************************")

	// Make a new slice big enough to hold elements of slice 1 and copy the
	// values over using the builtin copy function.
	slice3 := make([]string, len(slice1))
	copy(slice3, slice1)
	inspectSlice(slice3)

}

// SlicesExample4 is a sample program to show how to grow a slice using the built-in function append
// and how append grows the capacity of the underlying array.
func SlicesExample4() {
	// Declare a nil slice of strings.
	var data []string // this will yield -> nil | 0 | 0
	// data := []string{} // this will yield -> * (pointer) | 0 | 0 - pointer points to the empty struct value (struct{})

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Append ~100k strings to the slice.
	for record := 1; record <= 1e5; record++ {

		// Use the built-in function append to add to the slice.
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {

			// Calculate the percent of change.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for capacity.
			lastCap = cap(data)

			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChg)
		}
	}
}

// SlicesExample5 is a sample program to show how one needs to be careful when appending
// to a slice when you have a reference to an element.
func SlicesExample5() {
	// Declare a slice of 3 users.
	users := make([]sliceUser, 3)

	// Share the sliceUser at index 1.
	shareUser := &users[1]

	// Add a like for the sliceUser that was shared.
	shareUser.likes++

	// Display the number of likes for all users.
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// Add a new sliceUser.
	users = append(users, sliceUser{})

	// Add another like for the sliceUser that was shared.
	shareUser.likes++

	// Display the number of likes for all users.
	fmt.Println("*************************")
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// Notice the last like has not been recorded.
}

/*
	https://blog.golang.org/strings
	Go source code is always UTF-8.
	A string holds arbitrary bytes.
	A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
	Those sequences represent Unicode code points, called runes.
	No guarantee is made in Go that characters in strings are normalized.
	----------------------------------------------------------------------------
	Multiple runes can represent different characters:
	The lower case grave-accented letter à is a character, and it's also a code
	point (U+00E0), but it has other representations.
	We can use the "combining" grave accent code point, U+0300, and attach it to
	the lower case letter a, U+0061, to create the same character à.
	In general, a character may be represented by a number of different sequences
	of code points (runes), and therefore different sequences of UTF-8 bytes.
*/

// SlicesExample6 is a sample program to show how strings have a UTF-8 encoded byte array.
func SlicesExample6() {
	// Declare a string with both chinese and english characters.
	s := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune.
	var buf [utf8.UTFMax]byte

	// Iterate over the string.
	for i, r := range s {

		// Capture the number of bytes for this rune.
		rl := utf8.RuneLen(r)

		// Calculate the slice offset for the bytes associated
		// with this rune.
		si := i + rl

		// Copy of rune from the string to our buffer.
		copy(buf[:], s[i:si])

		// Display the details.
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}

// SlicesExample7 is a sample program to show how to declare and use variadic functions.
func SlicesExample7() {
	// display can accept and display multiple values of sliceUser types.
	display := func(users ...sliceUser) {
		fmt.Println("**************************")
		for _, u := range users {
			fmt.Printf("%+v\n", u)
		}
	}

	// change shows how the backing array is shared.
	change := func(users ...sliceUser) {
		users[1] = sliceUser{99, "Same Backing Array", 0}
	}

	// Declare and initialize a value of type sliceUser.
	u1 := sliceUser{
		id:   1432,
		name: "Betty",
	}

	// Declare and initialize a value of type sliceUser.
	u2 := sliceUser{
		id:   4367,
		name: "Janet",
	}

	// Display both sliceUser values.
	display(u1, u2)

	// Create a slice of sliceUser values.
	u3 := []sliceUser{
		{24, "Bill", 0},
		{32, "Joan", 0},
	}

	// Display all the sliceUser values from the slice.
	display(u3...)

	change(u3...)
	fmt.Println("**************************")
	for _, u := range u3 {
		fmt.Printf("%+v\n", u)
	}
}

// SlicesExample8 is a sample program to show how the for range has both value and pointer semantics.
func SlicesExample8() {
	// Using the value semantic form of the for range.
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", v)
	}

	fmt.Print("\n\n")

	// Using the pointer semantic form of the for range.
	friends = []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", friends[i])
	}
}

// SlicesAdvancedExample1 is a sample program to show how to use a third index slice.
func SlicesAdvancedExample1() {
	// Create a slice of strings with different types of fruit.
	slice := []string{"Apple", "Orange", "Banana", "Grape", "Plum"}
	inspectSlice(slice)

	// Take a slice of slice. We want just index 2
	takeOne := slice[2:3]
	inspectSlice(takeOne)

	// Take a slice of just index 2 with a length and capacity of 1
	takeOneCapOne := slice[2:3:3] // Use the third index position to
	inspectSlice(takeOneCapOne)   // set the capacity to 1.

	// Append a new element which will create a new
	// underlying array to increase capacity.
	takeOneCapOne = append(takeOneCapOne, "Kiwi")
	inspectSlice(takeOneCapOne)
}

// SlicesExercise1 is a to declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
func SlicesExercise1() {
	// Declare a nil slice of integers.
	var numbers []int

	// Append numbers to the slice.
	numbers = append(numbers, []int{1, 2, 3, 4, 5}...)

	// Display each value in the slice.
	for i, n := range numbers {
		fmt.Printf("%d: %v\n", i, n)
	}

	fmt.Print("\n\n")

	// Declare a slice of strings and populate the slice with names.
	names := []string{"James", "Jack", "Joe", "Mark"}

	// Display each index position and slice value.
	for i, n := range names {
		fmt.Printf("%d: %v\n", i, n)
	}

	fmt.Print("\n\n")

	// Take a slice of index 1 and 2 of the slice of strings.
	coolGuys := names[1:3:3]
	inspectSlice(coolGuys)

	// Display each index position and slice values for the new slice.
	for i, n := range coolGuys {
		fmt.Printf("Cool %d: %v\n", i, n)
	}
}
