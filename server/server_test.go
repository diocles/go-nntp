package nntpserver

import (
	"math"
	"testing"
)

func TestRangeEmpty(t *testing.T) {
	var rangeExpectations = []struct {
		input string
		low   int64
		high  int64
	}{
		{"", 0, math.MaxInt64},
		{"73-", 73, math.MaxInt64},
		{"73-1845", 73, 1845},
	}
	for _, e := range rangeExpectations {
		l, h := parseRange(e.input)
		if l != e.low {
			t.Fatalf("Error parsing %q, got low=%v, wanted %v",
				e.input, l, e.low)
		}
		if h != e.high {
			t.Fatalf("Error parsing %q, got high=%v, wanted %v",
				e.input, h, e.high)
		}
	}
}
