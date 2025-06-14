package main

import "math"

func minSizeSubarray(nums []int, target int) int {
	left := 0
	cnt := 0
	ans := math.MaxInt // 无限数组，求最小的子串长度，就应该找子串可能的最大值，这里设为∞

	sum := 0
	for _, v := range nums {
		sum += v
	}

	arrayCounts := 0 //target有可能是nums整个和的很多倍，也可能只在一个nums内就能找到最短ans。下面将这两者的逻辑统一起来了，就能够同一套代码解决整个问题。

	if target >= sum {
		arrayCounts = target / sum
		target = target % sum
	}

	// 下面的循环逻辑，仍然是先找到第一个合格的右端点，然后再通过左边进行删除以减到最小。
	l := len(nums)
	for right := range 2 * l { // 用一个实际的数组就能够完成n个数组的循环遍历
		cnt += nums[right%l]

		for cnt >= target {
			if cnt == target {
				ans = min(ans, right-left+1)
			}
			cnt -= nums[left%l]
			left++
		}
	}

	if ans == math.MaxInt { // 如果从头到尾ans都没有发生变化，那么说明一个都没有找到，就返回“没找到”
		return -1
	}

	return ans + arrayCounts*l // 记得返回结果要加上前面去掉的多余部分。
}

func minSizeSubarray1(nums []int, target int) int {
	left := 0
	cnt := 0
	minL := math.MaxInt

	sum := 0
	for _, v := range nums {
		sum += v
	}

	arrCounts := 0
	if target == sum {
		return len(nums)
	} else if sum < target {

	} else {
		arrCounts = target / sum
		target = target % sum
	}

	l := len(nums)
	for right := range 2 * l {
		cnt += nums[right%l]

		for cnt >= target {
			if cnt == target {
				minL = min(minL, right-left+1)
			}
			cnt -= nums[left]
			left++
		}
	}

	if minL == math.MaxInt {
		return -1
	}

	return minL + arrCounts*l

}
