//https://leetcode.cn/problems/fibonacci-number/submissions/
//dp方法，降低空间复杂度

package main

import "fmt"

func fib3(n int) int {
	if n<2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}

	return dp[1]
}

func main()  {
	v := fib3(10)
	fmt.Println(v)
}