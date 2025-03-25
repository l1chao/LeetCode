package doublepointer

func moveZeroes(nums []int) {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			for j = i + 1; nums[j] == 0; j++ {
			}
			nums[i] = nums[j]
			i++
		} else {

		}
	}
}
