package mymap

import (
	"fmt"
	"testing"
)

// 反转映射
// • 任务：编写函数，接收 map[string]int，返回反转后的 map[int]string（需处理值重复的情况，保留最后一个键）。
// • 示例输入：map["a":1, "b":2, "c":2]
// • 示例输出：map[1:"a" 2:"c"]

func ans4(map1 map[string]int) map[int]string {
	map2 := make(map[int]string)
	for k, v := range map1 {
		map2[v] = k
	}

	return map2
}

func Test4(t *testing.T) {
	fmt.Println(ans4(map[string]int{"a": 1, "b": 2, "c": 2}))
}
