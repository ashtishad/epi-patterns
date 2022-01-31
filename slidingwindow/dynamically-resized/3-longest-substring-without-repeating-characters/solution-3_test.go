package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		input  string
		result int
	}{
		{"abba", 2},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"dvdf", 3},
		{"abcdefghijklmnopqrstuvwxyz", 26},
	}

	for _, tc := range testCases {
		var want = tc.result
		got := lengthOfLongestSubstring(tc.input)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %v\n Expected: %v\n   Actual: %v\n",
				tc.input,
				want,
				got)
		}
	}
}
