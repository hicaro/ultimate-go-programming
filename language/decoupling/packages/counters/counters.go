// Package counters provides alert counter support.
package counters

// AlertCounter is an exported named type that
// contains an integer counter for alerts.
type AlertCounter int

// alertCounter is an unexported named type that
// contains an integer counter for alerts.
type alertCounter int

// New creates and returns values of the unexported type alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}
