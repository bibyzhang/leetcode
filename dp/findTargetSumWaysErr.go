//错误思路，先保留着吧

package main

import (
	"fmt"
	"math"
)

func findTargetSumWays(nums []int, target int) int {
	//-1000 <= target <= 1000，那么背包容量为-1000到1000,还有0
	//负数容量实际上为容量的最大价值存储标识为负数即可
	dp := make([]int, 2001)
	dp[0] = 0

	for i:=0; i<len(nums); i++ {
		for j:=1000; j>0; j-- {
			dp[j] = ftmax(dp[j], dp[j-nums[i]] + nums[i])
		}
	}

	times := 0
	for i:=-1000; i<=1000; i++ {
		bi := i
		if(i<0+target) {
			bi = int(math.Abs(float64(bi)))
			if dp[bi] == bi && (dp[bi+target] == bi+target) {
				fmt.Println(bi, dp[bi+target])
				times++
			}
		} else {
			if dp[i] == i && (dp[i-target] == i-target) {
				fmt.Println(i, dp[i-target])
				times++
			}
		}
	}

	return times
}

func ftmax(a int, b int) int {
	if a>b {
		return a
	}
	return b
}

func main()  {
	nums := []int{1,1,1,1,1}
	target := 3
	//nums := []int{1}
	//target := 1
	//nums := []int{1,0}
	//target := 1
	v := findTargetSumWays(nums, target)
	fmt.Println(v)
}