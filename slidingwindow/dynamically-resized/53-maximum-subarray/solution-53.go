package slidingwindow

import "math"

func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+sum < nums[i] {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		// update max
		max = int(math.Max(float64(sum), float64(max)))
	}

	return max
}
