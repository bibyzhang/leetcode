//https://leetcode.cn/problems/partition-equal-subset-sum/
package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	//不能整除，没有分半的和可以供两个子数组想加
	if sum%2 == 1 {
		return false
	}

	target := sum/2

	//题意，数值在1到100，数量为1到200，则sum的最大值为100*200=20000，则一半的容量最大为10000
	dp := make([]int, 10001)

	//初始化
	//dp[0] = 0

	for i:=0; i<len(nums); i++ {
		for j:=target; j>=nums[i]; j-- {
			dp[j] = cmax(dp[j], dp[j - nums[i]] + nums[i])
		}
	}

	if dp[target] == target {
		return true
	}

	return false
}

func cmax(a int, b int) int {
	if a>b {
		return a
	}

	return b
}

func main()  {
	//nums := []int{1, 5, 11, 5}
	nums := []int{1, 2, 3, 5}
	ok := canPartition(nums)
	fmt.Println(ok)
}