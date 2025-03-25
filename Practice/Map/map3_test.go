package mymap

// 键存在性检查
// • 任务：编写函数，接收一个 map[string]int 和键 key，返回该键的值；若键不存在，返回 0 并添加键值对 key:0。
// • 示例输入：map["a":1, "b":2], key="c"
// • 示例输出：0，且 map 变为 map[a:1 b:2 c:0]。

func ans3(hashMap map[string]int, key string) int {
	if v, ok := hashMap[key]; ok {
		return v
	}
	hashMap[key] = 0
	return 0

}
