package main

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	spaces := make([]int, 0)
	spaces = append(spaces, startTime[0])

	for i := 1; i < len(startTime); i++ {
		spaces = append(spaces, startTime[i]-endTime[i-1])
	}

	spaces = append(spaces, eventTime-endTime[len(endTime)-1])

	stat := 0
	maxN := 0

	for i := range spaces {
		stat += spaces[i]
		if i < k {
			continue
		}

		maxN = max(maxN, stat)

		stat -= spaces[i-(k+1)+1]
	}

	return maxN
}
