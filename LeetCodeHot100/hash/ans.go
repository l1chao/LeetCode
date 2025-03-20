package main

// https://leetcode.cn/problems/two-sum?envType=study-plan-v2&envId=top-100-liked

// 方法1，暴力遍历。

func main() {
	nums := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if fn(nums[i], nums[j]) {
				// ...
			}
		}
	}

	flag := false
	for _, v := range nums {
		if v == 100 { // 有，则从这里出去
			flag = true
			// return ...
		}
	}

	if !flag { // 如果没有，就从这里出去
		//...
	}
}

// 方法2，空间换时间
// 数组解决集合问题（有序容器找无序）：用hash。hash的key可以实现“立即找”。
func twoSum1(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := range nums {
		if v, ok := hashMap[target-nums[i]]; ok {
			return []int{v, i}
		}
		hashMap[nums[i]] = i
	}
	return nil

}
