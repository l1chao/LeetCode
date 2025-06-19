package main

import "fmt"

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
	
}

func main() {
	str := []int{1, 2, 3}
	is_you_a_silly_B := 1
	if str[0] == 1 {
		fmt.Println("nihao")
	}
}
