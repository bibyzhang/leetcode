package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]

	//第一步需要费用,到了i,是上一步费用加i费用，已经付完，可以走1-2步了
	for i:=2; i<len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]//选择费用低的到当前台阶
	}

	//i已经收费,到了最后两格楼梯，付完费都可以到楼顶了（一步或两步），选择费用低的
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func main()  {
	cost := []int{10,15,20}
	//cost := []int{1,100,1,1,1,100,1,1,100,1}
	c := minCostClimbingStairs(cost)
	fmt.Println(c)
}