//https://leetcode.cn/problems/fibonacci-number/submissions/
//dp方法

package main

import "fmt"

func fib2(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	v := fib2(10)
	fmt.Println(v)
}
