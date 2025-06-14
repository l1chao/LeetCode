package main

import "math"

func maxScore(cardPoints []int, k int) int {
	winLen := len(cardPoints) - k

	stat := 0
	minN := math.MaxInt
	total := 0
	for _, v := range cardPoints {
		total += v
	}

	if winLen == 0 {
		return total
	}

	for i, v := range cardPoints {
		stat += v
		if i < winLen-1 {
			continue
		}

		minN = min(minN, stat)

		stat -= cardPoints[i-winLen+1]
	}

	return total - minN
}
