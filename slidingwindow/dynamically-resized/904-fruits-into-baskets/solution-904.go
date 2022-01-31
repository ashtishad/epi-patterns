package slidingwindow

import "math"

func totalFruit(trees []int) int {
	baskets := make(map[int]bool)
	var max, s int

	for e, v := range trees {
		// basket available and different tree type
		if len(baskets) < 2 && !baskets[v] {
			baskets[v] = true
			max = int(math.Max(float64(max), float64(e-s+1)))
			continue
		}
		// collect fruits into existing basket
		if baskets[v] {
			max = int(math.Max(float64(max), float64(e-s+1)))
			continue
		}

		// none of previous two conditions are met
		baskets = make(map[int]bool) // reset baskets
		baskets[trees[e-1]] = true
		baskets[v] = true
		s = e - 1

		// checking if the last two trees are the same
		for trees[s] == trees[s-1] {
			s--
		}

		max = int(math.Max(float64(max), float64(e-s+1)))
	}
	return max
}
