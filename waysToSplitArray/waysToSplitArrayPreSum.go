//https://leetcode.cn/problems/number-of-ways-to-split-array/
//前缀和，先计算所有数综合，第二轮遍历叠加计算前面的值，总和减去这部分值就是后面额总和

package main

import "fmt"

func waysToSplitArrayPreSumn(nums []int) int {
	sum := 0
	for i:=0; i<len(nums); i++ {
		sum += nums[i]
	}

	c := 0
	pre_sum := 0
	//i右边至少有一个元素
	for i:=0; i<len(nums)-1; i++{
		pre_sum += nums[i]

		if(pre_sum>=(sum-pre_sum)) {
			c++
		}
	}

	return c
}

func main()  {
	//nums := []int{10,4,-8,7}
	nums := []int{2,3,1,0}
	//nums := []int{0, 0}
	//nums := []int{100000,-100000,100000,-100000}
	n := waysToSplitArrayPreSumn(nums)
	fmt.Println(n)
}