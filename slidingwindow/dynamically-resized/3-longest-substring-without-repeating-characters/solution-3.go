package slidingwindow

import "math"

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	var start, res int

	for end := range s {
		// if seen, restart the window just after last time it was seen
		if idx, ok := seen[s[end]]; ok {
			start = int(math.Max(float64(start), float64(idx+1)))
		}

		seen[s[end]] = end
		res = int(math.Max(float64(res), float64(end-start+1)))
	}

	return res
}
