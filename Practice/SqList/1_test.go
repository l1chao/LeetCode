package sqlist

// 删除无序数组的某元素x（交换元素到末尾 or 快慢指针）

func delete_x(arr []int, x int) {
	k := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			k++
		} else {
			arr[i-k] = arr[i]

		}
	}
}
