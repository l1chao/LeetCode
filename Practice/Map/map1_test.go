package mymap

import (
	"fmt"
	"testing"
)

// ​统计字符频率
// ​任务：编写函数，接收一个字符串，返回每个字符（包括空格、符号）出现的次数。
// ​示例输入："hello, world!"
// ​示例输出：map[h:1 e:1 l:3 o:2 ,:1 :1 w:1 r:1 d:1 !:1]
// ​扩展：忽略大小写（如 H 和 h 视为相同）。

func ans1(str string) map[string]int {
	hashMap := make(map[string]int)

	// 遍历都是给的键值对
	// 遍历字符串的字符。如果不需要序号，直接"_,v"
	for _, v := range str {
		hashMap[string(v)]++
	}

	return hashMap
}

func Test1(t *testing.T) {
	hashMap := ans1("hello, world!")
	fmt.Printf("%+v", hashMap)
}
