package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		res  float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
		{[]int{5}, 1, 5.00000},
	}

	for _, tc := range testCases {
		var want = tc.res
		got := findMaxAverage(tc.nums, tc.k)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %d\n Expected: %f\n   Actual: %f\n",
				tc.nums,
				want,
				got)
		}
	}
}
