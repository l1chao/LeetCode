package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	cnt := 0
	left := 0
	minL := math.MaxInt

	for right := range nums {
		cnt += nums[right]
		if cnt < target {
			continue
		}

		// for left < right && cnt-nums[left] >= target {
		// 	cnt -= nums[left]
		// 	left++
		// }
		for cnt-nums[left] >= target { // 实际上不需要left<right这一个条件，后面那个条件会包含这个条件
			cnt -= nums[left]
			left++
		}

		minL = min(minL, right-left+1)
	}
	if cnt < target {
		return 0
	}
	return minL
}

// target = 7, nums = [2,3,1,2,4,3]
func minSubArrayLen1(target int, nums []int) int {
	left := 0
	cnt := 0
	minL := len(nums)

	for i := range nums {
		cnt += nums[i]
		if cnt < target {
			continue
		}

		for cnt-nums[left] >= target {
			cnt -= nums[left]
			left++
		}

		minL = min(minL, i-left+1)
	}
	if cnt < target {
		return 0
	}
	return minL
}
