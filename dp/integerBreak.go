//https://leetcode.cn/problems/integer-break/

package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1

	for i:=3; i<=n; i++ {
		for j:=1; j<i-1; j++ {//i可以拆分未1到i-1个数字，比如4可以拆分为1，2，3,边界为j<i-1,因为j=i-1的时候i-j=i-i+1=1,1做乘法没意义，边界到j<i也没影响，只是多了五音
			//max(dp[i], xxx),取每一轮dp[i],即跟不同的j的乘积的最大值
			//3=(3-1)*1 = (3-2)*2 拆分两个数字时的拆法 (i-j)+j=i
			//dp[i-j]是拆分为多个数的场景,dp[i-j]即上i-j的最佳拆分(最大乘积)，除了初始dp[2]=1,其他依靠推导得出
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
	m := 10
	ma := integerBreak(m)
	fmt.Println(ma)
}