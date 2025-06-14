package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchInsert1(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
