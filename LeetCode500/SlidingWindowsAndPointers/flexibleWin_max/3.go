package main

// func lengthOfLongestSubstring(s string) int {
// 	dict := make(map[byte]int)
// 	stat := ""
// 	maxN := 0

// 	for i, v := range s {
// 		if _, ok := dict[s[i]]; !ok {
// 			dict[s[i]]++
// 			stat += string(v)
// 			maxN = max(maxN, len(stat))
// 		} else {
// 			stat += string(v)
// 			for dict[s[i]] > 1 {
// 				dict[stat[0]]--
// 				stat = stat[1:]
// 			}
// 		}
// 	}
// 	return maxN
// }

func lengthOfLongestSubstring(s string) int {
	dict := make(map[byte]int)
	left := 0
	maxN := 0

	for i := range s { // 对于每一个i：加-去重-比较最大值。
		dict[s[i]]++
		for dict[s[i]] > 1 {
			dict[s[left]]--
			left++
		}
		maxN = max(maxN, i-left+1)
	}
	return maxN
}

func lengthOfLongestSubstring1(s string) int {
	dict := make(map[byte]int)
	left := 0
	maxN := 0

	for i := range s {
		dict[s[i]]++
		for dict[s[i]] > 1 {
			dict[s[left]]--
			left++
		}
		maxN = max(maxN, i-left+1)
	}
	return maxN
}
