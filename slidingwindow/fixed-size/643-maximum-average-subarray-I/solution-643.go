// Package slidingwindow is a collection of solutions to the sliding window problem.
// 643. Maximum Average Subarray I - Easy
package slidingwindow

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var sum, s int
	res := math.Inf(-1) // -Inf

	for e := 0; e < len(nums); e++ {
		sum += nums[e]

		if (e - s + 1) == k {
			res = math.Max(res, float64(sum)/float64(k))
			// next window
			sum -= nums[s]
			s++
		}
	}
	return res
}

// Gotchas
// end - start +1 is the fixed window size.
// shifting next window is like just extract nums[start] from the window.
