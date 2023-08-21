package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for key, val := range nums {
		//每一个循环将当前值存起来
		//判断跟当前值相加能满足条件的值是否在前面出现(hashTable存起来则出现过)，如果，拿出其key
		if key_, ok := hashTable[target-val]; ok {
			return []int{key_, key}
		}
		//没找到则没匹配到，存储起来
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
