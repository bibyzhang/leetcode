//https://leetcode.cn/problems/remove-element
//双指针(双向指针)

package main

import "fmt"

func removeElement3(nums []int, val int) int {
	left,right := 0,len(nums)-1

	for left<=right {
		//找到左指针等于val的下标left
		//执行完循环left值就是等于val的下标
		for left<=right && nums[left]!=val {
			left++
		}
		//找到右指针不等于val的下标right
		//执行完训话right就是不等于val的下标
		for left<=right && nums[right]==val {
			right--
		}
		//把右边不等于val的值覆盖到左边等于val的值
		//left等于right其实就是同个值了,不需要再覆盖
		if left<right {
			nums[left] = nums[right]
			left++//从下一个开始村换
			right--//循环到上一个即可
		}
	}

	fmt.Println(nums)
	return left
}

func main()  {
	//nums := []int{3,2,2,3}
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	res := removeElement3(nums, val)
	fmt.Println(res)
}