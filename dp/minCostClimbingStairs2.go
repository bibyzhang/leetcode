package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i:=2; i<len(cost); i++ {
		if cost[i-1]>=cost[i-2] {
			dp[i] = dp[i-2] + cost[i]
		} else {
			dp[i] = dp[i-1] + cost[i]
		}
	}

	if(dp[len(cost)-1]>dp[len(cost)-2]) {
		return dp[len(cost)-2]
	}
	return dp[len(cost)-1]
}

func main()  {
	//cost := []int{10,15,20}
	cost := []int{1,100,1,1,1,100,1,1,100,1}
	c := minCostClimbingStairs(cost)
	fmt.Println(c)
}