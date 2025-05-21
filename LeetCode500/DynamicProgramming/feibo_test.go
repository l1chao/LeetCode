package dynamicprogramming

import (
	"fmt"
	"testing"
)

// 0 1 1 2 3 5 8 13 21 34 55
// 从0开始。且参数n>1
func fibo(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func fibo1(n int) int {
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	sum := 0
	for i := 2; i <= n; i++ {
		sum = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}

	return sum
}

func Test1(t *testing.T) {
	fmt.Println(fibo1(8))
}
