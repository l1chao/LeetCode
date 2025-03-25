package main

import "fmt"

func main() {
	set := make(map[int]string)
	set[1] = "true"
	k, v := set[2] // 如果没有，则前者是空值，后者是是否存在的bool
	fmt.Printf("k:[%v]v:[%v]", k, v)

	_, ok := set[1]
	fmt.Println(ok)
	if ok {

	}

	// 如果set表里面，key=1存在，则执行fn函数
	count := 100
	if _, ok := set[count]; ok {
		fn()
	}

	// 遍历set表,并判断
	for k, v := range set {
		fmt.Println(k)
		fmt.Println(v)
		fn(k, v)
	}
}

func fn(strs ...any) {
	if len(strs) == 0 {
		fmt.Println("执行fn函数")
	}
}
