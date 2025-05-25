package main

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	sum := 0
	scores := 0
	for i := range calories {
		sum += calories[i]
		if i < k-1 {
			continue
		}

		if sum > upper {
			scores++
		} else if sum < lower {
			scores--
		}

		sum -= calories[i-k+1]
	}

	return scores
}
