package datastructures

import "fmt"

// ArraysExample1 is a sample program to show how to declare and iterate over
// arrays of different types.
func ArraysExample1() {
	// Declare an array of five strings that is initialized
	// to its zero value.
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// Iterate over the array of strings.
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// Declare an array of 4 integers that is initialized
	// with some values.
	numbers := [4]int{10, 20, 30, 40}

	// Iterate over the array of numbers.
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}

// ArraysExample2 is a sample program to show how arrays of different sizes are
// not of the same type.
func ArraysExample2() {
	// Declare an array of 5 integers that is initialized
	// to its zero value.
	var five [5]int

	// Declare an array of 4 integers that is initialized
	// with some values.
	four := [4]int{10, 20, 30, 40}

	// Assign one array to the other
	// five = four

	// ./example2.go:21: cannot use four (type [4]int) as type [5]int in assignment

	fmt.Println(four)
	fmt.Println(five)
}

// ArraysExample3 is a sample program to show how the behavior of the for range and
// how memory for an array is contiguous.
func ArraysExample3() {
	// Declare an array of 5 strings initialized with values.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	// Iterate over the array displaying the value and
	// address of each element.
	for i, v := range friends {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &friends[i])
	}
}

// ArraysExample4 is a sample program to show how the for range has both value and pointer semantics.
func ArraysExample4() {
	// Using the pointer semantic form of the for range.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("Aft[%s]\n", friends[1])
		}
	}

	// Using the value semantic form of the for range.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}

	// Using the value semantic form of the for range but with pointer
	// semantic access. DON'T DO THIS.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range &friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}
}

// ArraysExercise1 is an exercise to:
// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
func ArraysExercise1() {
	// Declare an array of 5 strings set to its zero value.
	var colors [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	preColors := [5]string{"red", "white", "blue", "black", "green"}

	// Assign the populated array to the array of zero values.
	colors = preColors

	// Iterate over the first array declared.
	for i, v := range colors {
		// Display the string value and address of each element.
		fmt.Printf("Value: [%s]\t\tAddress: [%v]\t\tIndex Address: [%v]\n", v, &v, &colors[i])
	}
}
