package main

func minimumRecolors(blocks string, k int) int {
	stat := 0
	minN := k

	for i, v := range blocks {
		if v == 'W' {
			stat++
		}
		if i < k-1 {
			continue
		}

		minN = min(minN, stat)

		if blocks[i-k+1] == 'W' {
			stat--
		}
	}

	return minN
}
