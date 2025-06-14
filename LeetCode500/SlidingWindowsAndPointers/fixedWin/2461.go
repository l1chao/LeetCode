package main

func maximumSubarraySum(nums []int, k int) int64 {
	var stat int64
	var maxN int64
	dict := make(map[int]bool)

	for i, v := range nums {
		stat += int64(v)
		dict[v] = true

		if i < k-1 {
			continue
		}

		if len(dict) == k {
			maxN = max(maxN, stat)
		}

		stat -= int64(nums[i-k+1])
		delete(dict, nums[i-k+1])
	}

	return maxN
}
