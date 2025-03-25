package main

// 100,4,200,1,3,2
func ans3(nums []int) int {
	hashMap := make(map[int]bool)

	// 遍历都是k+v
	// 观点理解：
	// 1. 将所有数组nums里面的数，存入hash的键里面
	// 2. 将数组nums变为集合nums，集合nums能够直接查nums元素在不在。
	for _, v := range nums {
		hashMap[v] = true
	}

	longestStrike := 0
	for num := range hashMap {
		
		if _, ok := hashMap[num-1]; ok { // 前面有其他数字，跳过。具体验证就是value+ok
			continue
		}

		//如果前面没有其他数字了，看接下来的数字
		count := num + 1
		currentStrike := 1

		for {
			if _, ok := hashMap[count]; ok {
				currentStrike++
				count++
			} else {
				break
			}
		}

		if longestStrike < currentStrike {
			longestStrike = currentStrike
		}
	}
	return longestStrike
}

// 注意，如果是分别执行了常数次n循环，那么复杂度为n。
// 如果外面n循环，里面还要n循环，那复杂度就是n^2。
