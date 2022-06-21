//https://leetcode.cn/problems/binary-tree-postorder-traversal/
//后序遍历 左->右->中
//迭代法

package main

import "fmt"

type TreeNodePostOrder struct {
	Val int
	Left *TreeNodePostOrder
	Right *TreeNodePostOrder
}

func postorderTraversalIteration(root *TreeNodePostOrder) []int {
	var result []int
	if root == nil {
		return result
	}
	var stack []*TreeNodePostOrder

	stack = append(stack, root)
	for len(stack)!=0 {
		root = stack[len(stack)-1]//栈顶处理元素
		stack = stack[:len(stack)-1]//已处理元素出栈
		result = append(result, root.Val)
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}

	//上面结果是中右左
	//反转该数组得到左右中
	//for i:=0; i<len(result); i++ {
	//	for j:=len(result)-1; j>=0; j-- {
	//		result[i], result[j] = result[j], result[i]
	//	}
	//}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func main()  {
	tn3 := &TreeNodePostOrder{3, nil, nil}
	tn2 := &TreeNodePostOrder{2, tn3, nil}
	root := &TreeNodePostOrder{1, nil, tn2}
	arr := postorderTraversalIteration(root)
	fmt.Println(arr)
}
