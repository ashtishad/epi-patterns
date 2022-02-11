package binarysearch

func twoSum(nums []int, target int) []int {
	var n = len(nums)
	if n == 2 {
		return []int{1, 2}
	}

	var l, r = 0, n - 1
	for l < r {
		var sum = nums[l] + nums[r]
		if sum == target {
			return []int{l + 1, r + 1} // idx are 1-based
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
