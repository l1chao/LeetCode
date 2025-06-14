package main

func longestSubarray(nums []int) int {
	cnt := 0
	left := 0
	maxN := 0
	for i := range nums {
		if nums[i] == 0 {
			cnt++
		}
		for cnt > 1 {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}

		maxN = max(maxN, i-left+1)
	}
	return maxN
}
