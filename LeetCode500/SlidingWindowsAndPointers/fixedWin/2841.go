package main

func maxSum(nums []int, m int, k int) int64 {
	dict := make(map[int]int)
	var stat int64
	var maxSum int64

	for i, v := range nums {
		dict[v]++
		stat += int64(v) // 应该就在这里进行最大值加减！

		if i < k-1 {
			continue
		}

		if len(dict) >= m {
			maxSum = max(maxSum, stat)
		}

		if dict[nums[i-k+1]] == 1 {
			delete(dict, nums[i-k+1])
		} else {
			dict[nums[i-k+1]]--
		}
		stat -= int64(nums[i-k+1])
	}
	return int64(maxSum)
}
