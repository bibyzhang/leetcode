//错误思路，先保留着吧

package main

import (
	"fmt"
)

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if(sum+target < 0) {
		return 0
	}

	if target>sum {
		return 0
	}

	if (target+sum) % 2 == 1 {
		return 0
	}

	bgSize := (target+sum)/2

	dp := make([]int, bgSize+1)
	dp[0] = 1

	for i:=0; i<len(nums); i++ {
		for j:=bgSize; j>=nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[bgSize]
}

func main()  {
	//nums := []int{1,1,1,1,1}
	//target := 3
	//nums := []int{1}
	//target := 1
	//nums := []int{1,0}
	//target := 1
	nums := []int{100}
	target := -200
	v := findTargetSumWays(nums, target)
	fmt.Println(v)
}