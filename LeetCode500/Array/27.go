package main

func removeElement(nums []int, val int) int {
	k := 0
	for i := range nums {
		if nums[i] == val {
			k++
		} else {
			nums[i-k] = nums[i]
		}
	}
	nums = nums[:len(nums)-k]
	return len(nums)
}
