package main

func equalSubstring(s string, t string, maxCost int) int {
	left := 0
	maxN := 0

	for i := range s {
		maxCost -= abs(s[i], t[i])

		if maxCost < 0 {
			maxCost += abs(s[left], t[left])
			left++
		}
		maxN = max(maxN, i-left+1)
	}

	return maxN
}

func abs(a, b byte) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}
