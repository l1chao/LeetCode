package main

// 先加满到k个，然后加一再减一。
func maxVowels(s string, k int) int {
	cnt := 0
	maxNum := 0
	for i, v := range s {
		// 增
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			cnt++
		}
		if i < k-1 {
			continue
		}

		// 检
		maxNum = max(maxNum, cnt)

		// 删
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			cnt--
		}
	}
	return maxNum
}

func maxVowels1(s string, k int) int {
	stat := 0
	maxN := 0

	for i, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			stat++
		}
		if i < k-1 {
			continue
		}

		maxN = max(maxN, stat)

		if s[i-k+1] == 'a' || s[i-k+1] == 'e' || s[i-k+1] == 'i' || s[i-k+1] == 'o' || s[i-k+1] == 'u' {
			stat--
		}
	}
	return maxN
}
