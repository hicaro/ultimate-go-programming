package language

import (
	"fmt"
	"sort"
)

type mapUser struct {
	name    string
	surname string
}

// users defines a set of users.
type mapUsers []mapUser

// player represents someone playing our game.
type mapPlayer struct {
	name  string
	score int
}

// MapsExample1 is a sample program to show how to initialize a map, write to
// it, then read and delete from it.
func MapsExample1() {
	// Declare and make a map that stores values
	// of type mapUser with a key of type string.
	users := make(map[string]mapUser)

	// Add key/value pairs to the map.
	users["Roy"] = mapUser{"Rob", "Roy"}
	users["Ford"] = mapUser{"Henry", "Ford"}
	users["Mouse"] = mapUser{"Mickey", "Mouse"}
	users["Jackson"] = mapUser{"Michael", "Jackson"}

	// Read the value at a specific key.
	mouse := users["Mouse"]

	fmt.Printf("%+v\n", mouse)

	// Replace the value at the Mouse key.
	users["Mouse"] = mapUser{"Jerry", "Mouse"}

	// Read the Mouse key again.
	fmt.Printf("%+v\n", users["Mouse"])

	// Delete the value at a specific key.
	delete(users, "Roy")

	// Check the length of the map. There are only 3 elements.
	fmt.Println(len(users))

	// It is safe to delete an absent key.
	delete(users, "Roy")

	fmt.Println("Goodbye.")
}

// MapsExample2 is a sample program to show how maps behave when you read an
// absent key.
func MapsExample2() {
	// Create a map to track scores for players in a game.
	scores := make(map[string]int)

	// Read the element at key "anna". It is absent so we get
	// the zero-value for this map's value type.
	score := scores["anna"]

	fmt.Println("Score:", score)

	// If we need to check for the presence of a key we use
	// a 2 variable assignment. The 2nd variable is a bool.
	score, ok := scores["anna"]

	fmt.Println("Score:", score, "Present:", ok)

	// We can leverage the zero-value behavior to write
	// convenient code like this:
	scores["anna"]++

	// Without this behavior we would have to code in a
	// defensive way like this:
	if n, ok := scores["anna"]; ok {
		scores["anna"] = n + 1
	} else {
		scores["anna"] = 1
	}

	score, ok = scores["anna"]
	fmt.Println("Score:", score, "Present:", ok)
}

// MapsExample3 is a sample program to show how only types that can have
// equality defined on them can be a map key.
func MapsExample3() {
	// Declare and make a map that uses a slice as the key.
	// u := make(map[mapUsers]int)

	// ./example3.go:22: invalid map key type users
	u := make(map[string]int)

	// Iterate over the map.
	for key, value := range u {
		fmt.Println(key, value)
	}
}

// MapsExample4 is a sample program to show how to declare, initialize and iterate
// over a map. Shows how iterating over a map is random.
func MapsExample4() {
	// Declare and initialize the map with values.
	users := map[string]mapUser{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Iterate over the map printing each key and value.
	for key, value := range users {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Iterate over the map printing just the keys.
	// Notice the results are different.
	for key := range users {
		fmt.Println(key)
	}
}

// MapsExample5 is a sample program to show how to walk through a map by
// alphabetical key order.
func MapsExample5() {
	// Declare and initialize the map with values.
	users := map[string]mapUser{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Pull the keys from the map.
	var keys []string
	for key := range users {
		keys = append(keys, key)
	}

	// Sort the keys alphabetically.
	sort.Strings(keys)

	// Walk through the keys and pull each value from the map.
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
}

// MapsExample6 is a sample program to show that you cannot take the address
// of an element in a map.
func MapsExample6() {
	// Declare a map with initial values using a map literal.
	players := map[string]mapPlayer{
		"anna":  {"Anna", 42},
		"jacob": {"Jacob", 21},
	}

	// Trying to take the address of a map element fails.
	// anna := &players["anna"]
	// anna.score++

	// ./example4.go:23:10: cannot take the address of players["anna"]

	// Trying to increment in place
	// players["anna"].score++
	// cannot assign to struct field players["anna"].score in map

	fmt.Printf("Score: %d\n", players["anna"].score)

	// Instead take the element, modify it, and put it back.
	player := players["anna"]
	player.score++
	players["anna"] = player

	fmt.Printf("Score: %d\n", players["anna"].score)
}

// MapsExample7 is a sample program to show how maps are reference types.
func MapsExample7() {
	// double finds the score for a specific player and
	// multiplies it by 2.
	double := func(scores map[string]int, player string) {
		scores[player] = scores[player] * 2
	}
	// Initialize a map with values.
	scores := map[string]int{
		"anna":  21,
		"jacob": 12,
	}

	// Pass the map to a function to perform some mutation.
	double(scores, "anna")

	// See the change is visible in our map.
	fmt.Println("Score:", scores["anna"])
}

// MapsExercise1 expects to declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs
func MapsExercise1() {
	// Declare and make a map of integer type values.
	ages := make(map[string]int)
	// or:
	// ages := map[string]int{}

	// Initialize some data into the map.
	ages["John"] = 45
	ages["Jame"] = 82
	ages["Joe"] = 51

	// Display each key/value pair.
	for key, value := range ages {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
