package main

import "fmt"

func twoSumViolence(nums []int, target int) []int {
	for slowIndex := 0; slowIndex < len(nums); slowIndex++ {
		for fastIndex := (slowIndex + 1); fastIndex < len(nums); fastIndex++ {
			if nums[slowIndex] == target-nums[fastIndex] {
				return []int{slowIndex, fastIndex}
			}
		}
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	//nums := []int{3, 2, 4}
	//nums := []int{3,3}
	target := 9
	res := twoSumViolence(nums, target)
	fmt.Println(res)
}
