package decoupling

import (
	"fmt"
	"log"
	"time"
)

// admin represents an admin user with privileges.
type admin struct {
	person user // NOT Embedding
	level  string
}

// EmbeddingExample1 is a sample program to show how what we are doing is NOT embedding
// a type but just using a type as a field.
func EmbeddingExample1() {
	// Create an admin user.
	ad := admin{
		person: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can access fields methods.
	ad.person.notify()
}

// *****************************************************************************

// admin represents an admin user with privileges.
type useradmin struct {
	user  // Embedded Type
	level string
}

// EmbeddingExample2 is a sample program to show how to embed a type into another type and
// the relationship between the inner and outer type.
func EmbeddingExample2() {
	// Create an admin user.
	ad := useradmin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()
}

// *****************************************************************************

// EmbeddingExample3 is a sample program to show how embedded types work with interfaces.
func EmbeddingExample3() {
	// Create an admin user.
	ad := useradmin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is "promoted" to the outer type.
	sendNotification(&ad)
}

// *****************************************************************************

// admin represents an admin user with privileges.
type superadmin struct {
	user  // Embedded Type
	level string
}

// notify implements a method notifies admins
// of different events.
func (a *superadmin) notify() {
	fmt.Printf("Sending Admin Email To %s<%s>\n",
		a.name,
		a.email)
}

// EmbeddingExample4 is a sample program to show what happens when the outer and inner
// type implement the same interface.
func EmbeddingExample4() {
	// Create an admin user.
	ad := superadmin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is NOT "promoted" to the outer type.
	sendNotification(&ad)

	// We can access the inner type's method directly.
	ad.user.notify()

	// The inner type's method is NOT promoted.
	ad.notify()
}

// *****************************************************************************

// Document is the core data model we are working with.
type Document struct {
	Key   string
	Title string
}

// ==================================================

// Feed is a type that knows how to fetch Documents.
type Feed struct{}

// Count tells how many documents are in the feed.
func (f *Feed) Count() int {
	return 42
}

// Fetch simulates looking up the document specified by key. It is slow.
func (f *Feed) Fetch(key string) (Document, error) {
	time.Sleep(time.Second)

	doc := Document{
		Key:   key,
		Title: "Title for " + key,
	}
	return doc, nil
}

// ==================================================

// FetchCounter is the behavior we depend on for our process function.
type FetchCounter interface {
	Fetch(key string) (Document, error)
	Count() int
}

func process(fc FetchCounter) {
	fmt.Printf("There are %d documents\n", fc.Count())

	keys := []string{"a", "a", "a", "b", "b", "b"}

	for _, key := range keys {
		doc, err := fc.Fetch(key)
		if err != nil {
			log.Printf("Could not fetch %s : %v", key, err)
			return
		}

		fmt.Printf("%s : %v\n", key, doc)
	}
}

// ==================================================

// CachingFeed keeps a local copy of Documents that have already been
// retrieved. It embeds Feed to get the Fetch and Count behavior but
// "overrides" Fetch to have its cache.
type CachingFeed struct {
	// TODO embed *Feed and add a field for a map[string]Document.
	Feed
	documents map[string]Document
}

// NewCachingFeed initalizes a CachingFeed for use.
func NewCachingFeed(f *Feed) *CachingFeed {

	// TODO create a CachingFeed with an initialized map and embedded feed.
	c := CachingFeed{
		Feed:      Feed{},
		documents: make(map[string]Document),
	}

	// Return its address.
	return &c
}

// Fetch calls the embedded type's Fetch method if the key is not cached.
func (cf *CachingFeed) Fetch(key string) (Document, error) {

	// TODO implement this method. Check the map field for the specified key and
	// return it if found. If it's not found, call the embedded types Fetch
	// method. Store the result in the map before returning it.
	if doc, ok := cf.documents[key]; ok {
		return doc, nil
	}

	doc, err := cf.Feed.Fetch(key)
	if err != nil {
		return Document{}, err
	}

	cf.documents[key] = doc
	return doc, nil
}

// ==================================================

// EmbeddingExercise1 a program  which defines a type Feed with two methods: Count and Fetch. Create a
// new type CachingFeed that embeds *Feed but overrides the Fetch method.
//
// The CachingFeed type should have a map of Documents to limit the number of
// calls to Feed.Fetch.
func EmbeddingExercise1() {
	fmt.Println("Using Feed directly")
	process(&Feed{})

	// Call process again with your CachingFeed.
	fmt.Println("Using CachingFeed")
	c := NewCachingFeed(&Feed{})
	process(c)
}
