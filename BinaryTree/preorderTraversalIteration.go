//https://leetcode.cn/problems/binary-tree-preorder-traversal/
//前序遍历 中->左->右
//迭代法

package main

import (
	"fmt"
)

type TreeNodeIteration struct {
	Val int
	Left *TreeNodeIteration
	Right *TreeNodeIteration
}

func preorderTraversalIteration(root *TreeNodeIteration) []int {
	var result []int
	if root == nil {
		return result
	}

	s := []*TreeNodeIteration{}
	s = append(s, root)
	for len(s) != 0 {
		root := s[len(s)-1]//中
		s = s[0 : len(s)-1]
		result = append(result, root.Val)
		if root.Right != nil {
			s = append(s, root.Right)
		}
		if root.Left != nil {
			s = append(s, root.Left)
		}
	}

	return result
}

func main()  {
	tn3 := &TreeNodeIteration{3, nil, nil}
	tn2 := &TreeNodeIteration{2, tn3, nil}
	root := &TreeNodeIteration{1, nil, tn2}
	arr := preorderTraversalIteration(root)
	fmt.Println(arr)
}
