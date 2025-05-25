package main

func decrypt(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}

	n := len(code)
	r := k + 1
	if k < 0 {
		r = n
		k = -k
	}
	ans := make([]int, n)
	s := 0
	for _, x := range code[r-k : r] {
		s += x
	}

	for i := range ans {
		ans[i] = s
		s += code[r%n] - code[(r-k)%n]
		r++
	}
	return ans
}
