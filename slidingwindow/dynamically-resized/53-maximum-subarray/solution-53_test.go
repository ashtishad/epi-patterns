package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		nums   []int
		result int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-1, -2}, -1},
	}

	for _, tc := range testCases {
		var want = tc.result
		got := maxSubArray(tc.nums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %d\n Expected: %d\n   Actual: %d\n",
				tc.nums,
				want,
				got)
		}
	}
}
