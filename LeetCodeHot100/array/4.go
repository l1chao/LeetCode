package array

func maxSubArray(nums []int) int {
	// 找到头部下标
	i, j := 0, 0
	for {
		for ; nums[i] < 0; i++ {
		}

		j = i + 1
		total := nums[i]
		for ; nums[j] < 0; j++ {
			total += nums[j]
		}
		if total < 0 {
			i = j
		} else {
			break
		}
	}

	i, j = len(nums)-1, len(nums)-1
	for {
		for ; nums[i] < 0; i-- {
		}
		j = i - 1
		total := nums[i]
		for ; nums[j] < 0; j-- {
			total += nums[j]
		}
		if total < 0 {
			i = j
		}
	}

}
