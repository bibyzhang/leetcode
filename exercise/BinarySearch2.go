//https://leetcode.cn/problems/binary-search/

package main

import (
	"fmt"
	"math"
)

func binaray_search(nums []int, target int, low int, high int) int {
	//这些if根据特殊特殊边界情况写的具体值判定,非抽象逻辑,很不优雅啊
	if(low==high) {
		if nums[low]== target {
			return low
		}
		return -1
	}
	if(low+1==high) {
		if nums[low]== target {
			return low
		}
		if nums[high]== target {
			return high
		}
		return -1
	}
	//if(low>high)
	//终止条件应该用这个,当low+1到等于high,还是要查找的目标时,说明找不到

	x := (low + high)/2
	mid := int(math.Floor(float64(x)))

	//mid := low + (high-low)/2 应该用这个取中间数
	//当到最后一个元素时,mid就是最后一个元素,我上面的写法到4,5的时候,取出来还是4

	if(target>nums[mid]) {
		low = mid //这里可以加1呀,mid已经不符合了,往后+1
		return binaray_search(nums, target, low, high)
	} else if(target<nums[mid]) {
		high = mid
		return binaray_search(nums, target, low, high)
	}

	return mid
}

func search1(nums []int, target int) int {
	low := 0
	high := len(nums)-1
	return binaray_search(nums, target, low, high)
}

func main() {
	nums := []int{5}
	target := -5
	value := search1(nums, target)
	fmt.Println(value)
}