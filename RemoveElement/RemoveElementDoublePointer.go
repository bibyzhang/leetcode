//https://leetcode.cn/problems/remove-element
//双指针(快慢指针)

package main

import "fmt"

func removeElement2(nums []int, val int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex<len(nums); fastIndex++ {
		if(nums[fastIndex]!=val) {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}

	fmt.Println(nums)
	return slowIndex
}

func main()  {
	//nums := []int{3,2,2,3}
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	res := removeElement2(nums, val)
	fmt.Println(res)
}