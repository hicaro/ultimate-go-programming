// Package users provides support for user management.
package users

// User represents information about a user.
type User struct {
	Name string
	ID   int

	password string
}

// user represents information about a user.
type user struct {
	Name string
	ID   int
}

// Manager represents information about a manager.
type Manager struct {
	Title string

	user
}
