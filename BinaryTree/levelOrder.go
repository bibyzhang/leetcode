//https://leetcode.cn/problems/binary-tree-level-order-traver
package main

import "fmt"

type TreeNodeLevel struct {
	Val int
	Left *TreeNodeLevel
	Right *TreeNodeLevel
}

func levelOrder(root *TreeNodeLevel) [][]int {
	var results [][]int
	var queue []*TreeNodeLevel

	if root == nil {
		return results
	}

	queue = append(queue, root)
	for len(queue)>0 {
		var result []int
		size := len(queue)
		for i:=0; i<size; i++ {
			root = queue[0]//从队列取值
			queue = queue[1:]//值出队列
			result = append(result, root.Val)
			if root.Left!=nil {
				queue = append(queue, root.Left)
			}
			if root.Right!=nil {
				queue = append(queue, root.Right)
			}
		}
		results = append(results, result)
	}

	return results
}

func main()  {
	tn5 := &TreeNodeLevel{7, nil, nil}
	tn4 := &TreeNodeLevel{15, nil, nil}
	tn3 := &TreeNodeLevel{20, tn4, tn5}
	tn2 := &TreeNodeLevel{9, nil, nil}
	root := &TreeNodeLevel{3, tn2, tn3}
	arr := levelOrder(root)
	fmt.Println(arr)
}