package main

func removeDuplicates(nums []int) int {
	slow, fast := 1, 1

	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {

}
