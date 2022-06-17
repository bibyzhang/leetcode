//https://leetcode.cn/problems/binary-tree-postorder-traversal/
//后序遍历 左->右->中
package main

import "fmt"

type TreeNode2 struct {
	Val int
	Left *TreeNode2
	Right *TreeNode2
}

func postorderTraversal(root *TreeNode2) []int {
	var arr []int
	traversal2(root, &arr)
	return arr
}

func traversal2(root *TreeNode2, arr *[]int)  {
	if root == nil {
		return
	}
	traversal2(root.Left, arr)
	traversal2(root.Right, arr)
	*arr = append(*arr, root.Val)
}

func main()  {
	t3 := &TreeNode2{3, nil, nil}
	t2 := &TreeNode2{2, t3, nil}
	root := &TreeNode2{1, nil, t2}
	r := postorderTraversal(root)
	fmt.Println(r)
}