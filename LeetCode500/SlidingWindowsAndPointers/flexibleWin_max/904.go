package main

func totalFruit(fruits []int) int {
	maxN := 0
	dict := make(map[int]int)
	left := 0

	for i := range fruits {
		dict[fruits[i]]++

		for len(dict) > 2 {
			dict[fruits[left]]--
			if dict[fruits[left]] == 0 {
				delete(dict, fruits[left])
			}
			left++
		}

		maxN = max(maxN, i-left+1)
	}
	return maxN
}
