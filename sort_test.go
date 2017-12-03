package dateparse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestSort tests Sort implementation
func TestSort(t *testing.T) {
	var times []time.Time
	var timesToSort []time.Time
	var timeStrSlice = []string{
		"2017-11-11 11:01:01",
		"2009-11-22T11:22:02",
		"2009-11-29T04:44:04-04:00",
		"2009-11-23T03:33:03",
		"",
	}
	for _, str := range timeStrSlice {
		if t, err := ParseAny(str); err == nil {
			timesToSort = append(timesToSort, t)
			times = append(times, t)
		}
	}

	Sort(timesToSort)
	t.Logf("original times: %+v\n", times)
	t.Logf("- sorted times: %+v\n", timesToSort)
	assert.Equal(t, times[1], timesToSort[0])
	assert.Equal(t, times[3], timesToSort[1])
	assert.Equal(t, times[2], timesToSort[2])
	assert.Equal(t, times[0], timesToSort[3])
}
