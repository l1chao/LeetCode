package main

// 输入：arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
// 输出：3
func numOfSubarrays(arr []int, k int, threshold int) int {
	cnt := 0
	total := 0

	for i, v := range arr {
		cnt += v
		if i < k-1 {
			continue
		}

		if cnt >= threshold*k {
			total++
		}

		cnt -= arr[i-k+1]
	}

	return total
}

func numOfSubarrays1(arr []int, k int, threshold int) int {
	stat := 0
	ans := 0

	for i, v := range arr {
		stat += v
		if i < k-1 {
			continue
		}

		if stat >= k*threshold {
			ans++
		}

		stat -= arr[i-k+1]
	}
	return ans
}
