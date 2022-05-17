//https://leetcode.cn/problems/climbing-stairs/
package main

import "fmt"

func climbStairs(n int) int {
	if n<3 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i:=3; i<=n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main()  {
	c := climbStairs(9)
	fmt.Println(c)
}