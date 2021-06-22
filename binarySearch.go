package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	return bSearch(nums, low, high, target)
}

func bSearch(nums []int, low, high, target int) int {
	if(low>high) {
		return -1
	}
	mid := low + (high-low)/2//向下取整
	if( nums[mid]>target ) {
		return bSearch(nums, low, mid-1, target)
	} else if( nums[mid]<target ) {
		return bSearch(nums, mid+1, high, target)
	}

	return mid
}

func main() {
	//nums := []int{-1,0,3,5,9,12}
	//nums := []int{1,2,3,4,5,6,7,8,9}
	nums := []int{-1,0,3,5,9,12}
	//target := 9
	//target := 5
	target := 2
	value := search(nums, target)
	fmt.Println(value)
}