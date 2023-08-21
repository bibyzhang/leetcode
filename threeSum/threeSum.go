// https://leetcode.cn/problems/3sum/
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	//先排序
	sort.Ints(nums)
	fmt.Println(nums)

	//存储满足条件的结果
	//var result [][]int
	result := [][]int{}

	//计算三数之和
	for i := 0; i < len(nums); i++ {
		//排序后第一个被计算的元素已经大于0
		//那么三元数组注定无法凑成，直接返回结果
		if nums[i] > 0 {
			return result
		}

		//排序后相同的值放在一起，如果前面已经有相等的值，说明已经以该值为第一个元素寻找过三元数组，同样的起始值会匹配到一样的结果
		//如果是后续值跟起始值一样是允许的，比如 [-1,-2,2]
		if i > 0 && (nums[i] == nums[i-1]) {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right { //left等于right为同一个元素，对三元数组来说没有意义
			if (nums[i] + nums[left] + nums[right]) > 0 {
				right--
			} else if (nums[i] + nums[left] + nums[right]) < 0 {
				left++
			} else { //找到结果,将数据存储到数组
				res_s := []int{nums[i], nums[left], nums[right]}
				result = append(result, res_s)

				//i值不变，right值相同，结果会相同
				for (left < right) && (nums[right] == nums[right-1]) {
					right--
					continue
				}

				//i值不变，left值相同，结果会相同
				for (left < right) && (nums[left] == nums[left+1]) {
					left++
					continue
				}

				//找到答案后，继续收缩指针寻找新三元组
				right--
				left++
			}
		}
	}

	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//排序后 [-4 -1 -1 0 1 2]
	res := threeSum(nums)
	fmt.Println(res)
}
