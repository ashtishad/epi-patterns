package binarysearch

import (
	"reflect"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
	}

	for _, tc := range testCases {
		var want = tc.result
		got := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("failed! got: %v, want: %v, input: %v, k: %v", got, want, tc.nums, tc.target)
		}
	}
}
