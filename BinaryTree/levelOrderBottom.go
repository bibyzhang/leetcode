//https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/

package main

import "fmt"

type TreeNodeLevelOrderBottom struct {
	Val int
	Left *TreeNodeLevelOrderBottom
	Right *TreeNodeLevelOrderBottom
}

func levelOrderBottom(root *TreeNodeLevelOrderBottom) [][]int {
	var results [][]int

	if root == nil {
		return results
	}

	var queue []*TreeNodeLevelOrderBottom
	queue = append(queue, root)

	for len(queue)>0 {
		size := len(queue)
		var result []int

		for i:=0; i<size; i++ {
			root = queue[0]
			queue = queue[1:]
			result = append(result, root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}

		results = append(results, result)
	}

	//反转
	for i,j := 0, len(results)-1; i<j; i,j = i+1,j-1 {
		results[i], results[j] = results[j], results[i]
	}

	return results
}

func main()  {
	tn5 := &TreeNodeLevelOrderBottom{7, nil, nil}
	tn4 := &TreeNodeLevelOrderBottom{15, nil, nil}
	tn3 := &TreeNodeLevelOrderBottom{20, tn4, tn5}
	tn2 := &TreeNodeLevelOrderBottom{9, nil, nil}
	root := &TreeNodeLevelOrderBottom{3, tn2, tn3}
	arr := levelOrderBottom(root)
	fmt.Println(arr)
}
