//https://leetcode.cn/problems/binary-tree-inorder-traversal/
//中序遍历 左 -> 中 -> 右
//迭代方法

package main

import "fmt"

type TreeNodeIterationInOrder struct {
	Val int
	Left *TreeNodeIterationInOrder
	Right *TreeNodeIterationInOrder
}

func inorderTraversalIteration(root *TreeNodeIterationInOrder) []int {
	var result []int
	var stack []*TreeNodeIterationInOrder
	for root!=nil || len(stack)!=0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1] //出栈,处理元素
			stack = stack[:len(stack)-1]
			result = append(result, root.Val)
			root = root.Right
		}
	}

	return result
}

func main()  {
	tn1 := &TreeNodeIterationInOrder{3, nil, nil}
	tn2 := &TreeNodeIterationInOrder{2, tn1, nil}
	root := &TreeNodeIterationInOrder{1, nil, tn2}
	arr := inorderTraversalIteration(root)
	fmt.Println(arr)
}