// Package toy contains support for managing toy inventory.
package toy

// Toy is a struct with four fields. Name string,
// Weight int, onHand int and sold int.
type Toy struct {
	Name   string
	Weight int
	onHand int
	sold   int
}

// New is a function that accepts values for the
// exported fields. Return a pointer of type Toy that is initialized
// with the parameters.
func New(name string, weight int) *Toy {
	t := Toy{
		Name:   name,
		Weight: weight,
	}

	return &t
}

// OnHand is a method with a pointer receiver that
// returns the current on hand count.
func (t *Toy) OnHand() int {
	return t.onHand
}

// UpdateOnHand is a method with a pointer receiver that
// updates and returns the current on hand count.
func (t *Toy) UpdateOnHand(n int) int {
	t.onHand = n
	return t.onHand
}

// Sold is a method with a pointer receiver that
// returns the current sold count.
func (t *Toy) Sold() int {
	return t.sold
}

// UpdateSold is a method with a pointer receiver that
// updates and returns the current sold count.
func (t *Toy) UpdateSold(n int) int {
	t.sold = n
	return t.sold
}
