package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		s   string
		res int
	}{
		{"QWER", 0},
		{"QQWE", 1},
		{"QQQW", 2},
	}

	for _, tc := range testCases {
		var want = tc.res
		got := balancedString(tc.s)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %s\n Expected: %d\n   Actual: %d\n",
				tc.s,
				want,
				got)
		}
	}
}
