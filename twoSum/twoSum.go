package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for key, val := range nums {
		if key_, ok := hashTable[target-val]; ok {
			return []int{key_, key}
		}
		hashTable[val] = key
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	//nums := []int{3, 2, 4}
	//nums := []int{3,3}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
