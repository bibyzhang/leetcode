//https://leetcode.cn/problems/integer-break/

package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1

	for i:=3; i<=n; i++ {
		for j:=1; j<i-1; j++ {//i可以拆分未1到i-1个数字，比如4可以拆分未1，2，3
			//max(dp[i], xxx),取每一轮dp[i],即跟不同的j的乘积的最大值
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
		}
	}

	return dp[n]
}

func max(a int, b int) int {
	if a>b {
		return a
	}

	return b
}

func main()  {
	m := 3
	ma := integerBreak(m)
	fmt.Println(ma)
}