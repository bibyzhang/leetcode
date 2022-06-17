//https://leetcode.cn/problems/binary-tree-inorder-traversal/
//中序遍历 左 -> 中 -> 右
package main

import "fmt"

type TreeNode3 struct {
	Val int
	Left *TreeNode3
	Right *TreeNode3
}

func inorderTraversal(root *TreeNode3) []int {
	var arr []int
	traversal3(root, &arr)
	return arr
}

func traversal3(root *TreeNode3, arr *[]int)  {
	if root==nil {
		return
	}
	traversal3(root.Left, arr)
	*arr = append(*arr, root.Val)
	traversal3(root.Right, arr)
}

func main()  {
	t3 := &TreeNode3{3, nil, nil}
	t2 := &TreeNode3{2, t3, nil}
	root := &TreeNode3{1, nil, t2}
	r := inorderTraversal(root)
	fmt.Println(r)
}