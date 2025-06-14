package main

func distinctNumbers(nums []int, k int) []int {
	ans := make([]int, 0)
	dict := make(map[int]int)

	for i := range nums {
		dict[nums[i]]++
		if i < k-1 {
			continue
		}

		ans = append(ans, len(dict))

		dict[nums[i-k+1]]--
		if dict[nums[i-k+1]] == 0 {
			delete(dict, nums[i-k+1])
		}
	}

	return ans
}
