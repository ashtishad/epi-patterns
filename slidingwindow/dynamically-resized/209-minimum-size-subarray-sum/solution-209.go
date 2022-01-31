package slidingwindow

func minSubArrayLen(target int, nums []int) int {
	var sum, s int
	n := len(nums)
	res := n + 1 // maximum length for problem context

	for e := 0; e < n; e++ {
		sum += nums[e]
		for sum >= target {
			res = min(res, e-s+1)
			sum -= nums[s]
			s++
		}
	}

	// check value updated or not
	if res == n+1 {
		res = 0
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
