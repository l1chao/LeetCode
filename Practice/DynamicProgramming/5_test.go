package dynamicprogramming

import (
	"fmt"
	"testing"
)

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func Test5(t *testing.T) {
	m := 3
	dp := make([][]int, m)
	fmt.Println(dp[0])
	// fmt.Println(dp[0][0])
	fmt.Println(len(dp[0]))
	fmt.Println(cap(dp[0]))
	dp[0] = append(dp[0], 9527)
	fmt.Println(dp[0][0])
	fmt.Printf("%#v\n", dp[0])
	fmt.Printf("%T\n", dp[0])
	fmt.Printf("%p\n", dp[0])
	fmt.Println()
	fmt.Println(uniquePaths(3, 7))
}
