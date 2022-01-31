package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		target int
		nums   []int
		result int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	for _, tc := range testCases {
		var want = tc.result
		got := minSubArrayLen(tc.target, tc.nums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %d\n Expected: %d\n   Actual: %d\n",
				tc.nums,
				want,
				got)
		}
	}
}
