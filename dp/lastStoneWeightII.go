package main

import (
	"fmt"
	"math"
)

func lastStoneWeightII(stones []int) int {
	sum := 0
	//重量和为背包容量，总和为需要背包的最大容量
	for _, v := range stones {
		sum += v
	}

	//向下取整为能达到的拆分最大容量
	target := int(math.Floor(float64(sum/2)))

	//30*1000/2 = 15000
	dp := make([]int, 15001)
	dp[0] = 0

	//遍历商品
	//遍历背包,反向
	for i:=0; i<len(stones); i++ {
		for j:=target; j>=stones[i]; j-- {
			dp[j] = lsmax(dp[j], dp[j-stones[i]] + stones[i])
		}
	}

	lsw := (sum-dp[target]) - dp[target]

	return lsw
}

func lsmax(a int, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	//stones := []int{2, 4, 1, 1}
	//stones := []int{2,7,4,1,8,1}
	stones := []int{31,26,33,21,40}
	v := lastStoneWeightII(stones)
	fmt.Println(v)
}