//https://leetcode.cn/problems/remove-element

package main

import "fmt"

func removeElement(nums []int, val int) int {
	arr_num := len(nums)
	for i := 0; i <= arr_num-1; i++ {
		if(nums[i]==val) {
			//集体往前移动
			for j := i; j < arr_num-1; j++ {
				nums[j] = nums[j+1]
			}
			arr_num--
			i--
		}
	}

	fmt.Println(nums)
	return arr_num
}

func main()  {
	//nums := []int{3,2,2,3}
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	res := removeElement(nums, val)
	fmt.Println(res)
}