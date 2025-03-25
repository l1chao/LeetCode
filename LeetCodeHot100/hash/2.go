package main

// 先写没有重复的情况
func ans2(strs []string) [][]string {

	// var i1, i2 interface{} = 10, 10

	result := make([][]string, 0)

	hashMap := make(map[uint32][]string)

	for _, str := range strs {
		var count uint32
		for _, v := range str {
			count = count | 2<<(byte(v)-'a')
		}
		hashMap[count] = append(hashMap[count], str)
	}

	for _, v := range hashMap {
		result = append(result, v)
	}
	return result
}

func ans21(strs []string) [][]string {

	hashMap := make(map[[26]int][]string)

	for _, v := range strs {
		count := [26]int{}
		for _, single := range []byte(v) {
			count[single-'a']++
		}
		hashMap[count] = append(hashMap[count], v)
	}

	result := make([][]string, 0)
	for _, v := range hashMap {
		result = append(result, v)
	}

	return result
}
