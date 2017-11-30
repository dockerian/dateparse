// Package dateparse - sort.go
package dateparse

import (
	"sort"
	"time"
)

// TimeSlice defines a slice of time.Time
type TimeSlice []time.Time

// Sort uses implemented sort.Interace by TimeSlice
func Sort(timeSlice []time.Time) {
	sort.Sort(TimeSlice(timeSlice))
}

// Len implements sort.Interface
func (ts TimeSlice) Len() int {
	return len(ts)
}

// Less implements sort.Interface
func (ts TimeSlice) Less(i, j int) bool {
	return ts[i].Before(ts[j])
}

// Swap implements sort.Interface
func (ts TimeSlice) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}
