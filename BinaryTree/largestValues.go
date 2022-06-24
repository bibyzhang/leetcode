//https://leetcode.cn/problems/find-largest-value-in-each-tree-row/

package main

import "fmt"

type TreeNodeLargest struct {
	Val int
	Left *TreeNodeLargest
	Right *TreeNodeLargest
}

func largestValues(root *TreeNodeLargest) []int {
	var result []int
	if root == nil {
		return result
	}

	var queue []*TreeNodeLargest
	queue = append(queue, root)
	for len(queue)>0 {
		size := len(queue)

		max := -2147483648
		for i:=0; i<size; i++ {
			root = queue[0]
			queue = queue[1:]
			//if root.Val>max || max == 0 {
			if root.Val>max {
				max = root.Val
			}
			if root.Left!=nil {
				queue = append(queue, root.Left)
			}
			if root.Right!=nil {
				queue = append(queue, root.Right)
			}
		}

		result = append(result, max)
	}

	return result
}

func main()  {
	//tn6 := &TreeNodeLargest{9, nil, nil}
	//tn5 := &TreeNodeLargest{3, nil, nil}
	//tn4 := &TreeNodeLargest{5, nil, nil}
	//tn3 := &TreeNodeLargest{2, nil, tn6}
	//tn2 := &TreeNodeLargest{3, tn4, tn5}
	//root := &TreeNodeLargest{1, tn2, tn3}

	tn1 := &TreeNodeLargest{-1, nil, nil}
	root := &TreeNodeLargest{0, tn1, nil}

	//tn2 := &TreeNodeLargest{-37, nil, nil}
	//tn1 := &TreeNodeLargest{0, nil, nil}
	//root := &TreeNodeLargest{40, tn1, tn2}
	arr := largestValues(root)
	fmt.Println(arr)
}