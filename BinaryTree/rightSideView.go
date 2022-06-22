//https://leetcode.cn/problems/binary-tree-right-side-view/

package main

import "fmt"

type TreeNodeRightSideView struct {
	Val int
	Left *TreeNodeRightSideView
	Right *TreeNodeRightSideView
}

func rightSideView(root *TreeNodeRightSideView) []int {
	var result []int

	if root == nil {
		return result
	}

	var queue []*TreeNodeRightSideView
	queue = append(queue, root)
	for len(queue)>0 {
		size := len(queue)
		for i:=0; i<size; i++ {
			root = queue[0]
			queue = queue[1:]
			if i == size-1 {
				result = append(result, root.Val)
			}
			if root.Left!=nil {
				queue = append(queue, root.Left)
			}
			if root.Right!=nil {
				queue = append(queue, root.Right)
			}
		}
	}

	return result
}

func main()  {
	tn5 := &TreeNodeRightSideView{4, nil, nil}
	tn4 := &TreeNodeRightSideView{5, nil, nil}
	tn3 := &TreeNodeRightSideView{3, nil, tn5}
	tn2 := &TreeNodeRightSideView{2, nil, tn4}
	root := &TreeNodeRightSideView{1, tn2, tn3}
	arr := rightSideView(root)
	fmt.Println(arr)
}