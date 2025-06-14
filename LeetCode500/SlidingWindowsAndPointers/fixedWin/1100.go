package main

func numKLenSubstrNoRepeats(s string, k int) int {
	dict := make(map[byte]int)
	ans := 0

	for i := range s {
		dict[s[i]]++
		if i < k-1 {
			continue
		}

		if len(dict) == k {
			ans++
		}

		dict[s[i-k+1]]--
		if dict[s[i-k+1]] == 0 {
			delete(dict, s[i-k+1])
		}
	}
	return ans
}
