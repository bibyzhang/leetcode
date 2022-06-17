package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var res []int
	traversal(root, &res)
	return res
}

func traversal(current *TreeNode, res *[]int)  {
	if current==nil {
		return
	}
	*res = append(*res, current.Val)
	traversal(current.Left, res)
	traversal(current.Right, res)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	t3 := &TreeNode{3, nil, nil}
	t2 := &TreeNode{2, t3, nil}
	root := &TreeNode{1, nil, t2}
	res := preorderTraversal(root)
	fmt.Println(res)
}