package main

import (
	"fmt"
	"testing"
)

// {2,3,1,1,44,42,32,1,91,1}
// 删除无序数组的某元素x，删除前后其余元素不要求保持顺序。
func deleteDeputy(arr *[]int, x int) {
	k := 0
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == x {
			k++
		} else {
			(*arr)[i-k] = (*arr)[i]
		}
	}
	// arr = &((*arr)[:len(arr)-k])
}

type Person struct {
	Name string
	Age  int
}

func TestMain(t *testing.T) {
	var p Person
	fmt.Printf("%#v\n", p)
	var arr []int
	fmt.Printf("%v\n", arr)
	fmt.Println(arr == nil)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	arr = append(arr, 1)
	fmt.Println(arr[0])

	arr1 := make([]int, 0)
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
	fmt.Println(arr1 == nil)

	var i, j int
	for ; j < 10; i, j = i+1, j+1 {

	}
}
