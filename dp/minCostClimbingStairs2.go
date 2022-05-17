//https://leetcode.cn/problems/min-cost-climbing-stairs/submissions/
//第一步不付费版本，即先付费在跳

package main

import "fmt"

func minCostClimbingStairs2(cost []int) int {
	dp := make([]int, len(cost)+1)//第一步不用付费，要走到楼顶才付了费，在当前楼梯不付费，要走的时候才付
	dp[0] = 0
	dp[1] = 0 //可以从第一格或第二格开始走
	//i等于cost的时候是楼顶
	for i:=2; i<=len(cost); i++ {
		dp[i] = min2(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])//到i要加上i-1或i-2的费用
	}
	return dp[len(cost)]
}

func min2(a int, b int) int {
	if a>b {
		return b
	}

	return a
}

func main()  {
	//cost := []int{10,15,20}
	cost := []int{1,100,1,1,1,100,1,1,100,1}
	c := minCostClimbingStairs2(cost)
	fmt.Println(c)
}