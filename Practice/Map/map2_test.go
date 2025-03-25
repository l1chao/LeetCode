package mymap

import (
	"fmt"
	"testing"
)

// ​去重切片
// ​任务：编写函数，接收一个整数切片，返回去重后的新切片，顺序不限。
// ​示例输入：[]int{3, 1, 4, 1, 5, 9, 2, 6, 5}
// ​示例输出：[]int{3, 1, 4, 5, 9, 2, 6}（顺序可不同）
// ​提示：用 map 记录已存在的元素。

func ans2(nums []int) []int {
	hashMap := make(map[int]bool)

	for _, v := range nums {
		hashMap[v] = true
	}

	fmt.Println("len of hashMap:", len(hashMap))
	result := make([]int, len(hashMap))
	i := 0
	for k := range hashMap {
		result[i] = k
		i++
	}

	return result
}

func Test2(t *testing.T) {
	fmt.Println(ans2([]int{3, 1, 4, 1, 5, 9, 2, 6, 5}))
}

// 首先应该放慢速度，在思维上将整个过程理顺，哪怕很慢。
// 然后组合起来，越来越加快速度。
// 全程应该保持注意力高度专注。在训练的时候，要耐心努力冲破自己的极限，不遗余力。但是结束之后也应该注意休息。

// 获取map所有的keys
func ans21(hashMap map[string]int) []any {
	keys := make([]any, 0, len(hashMap))
	for k := range hashMap {
		keys = append(keys, k)
	}
	return keys
}
