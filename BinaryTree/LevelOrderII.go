package main

import (
	"fmt"
)

type Node struct {
	Val int
	Children []*Node
}

func LevelOrderII(root *Node) [][]int {
	var results [][]int

	if root == nil {
		return results
	}

	var queue []*Node
	queue = append(queue, root)
	for len(queue)>0 {
		var result []int
		size := len(queue)

		for i:=0; i<size; i++ {
			root = queue[0]
			queue = queue[1:]
			result = append(result, root.Val)
			if len(root.Children)>0 {
				for _, val := range root.Children {
					if val!=nil {
						queue = append(queue, val)
					}
				}
			}
		}

		results = append(results, result)
	}

	return results
}

func main() {
	node := []*Node{nil}
	n6 := &Node{6, node}
	n5 := &Node{5, node}
	n4 := &Node{4, node}
	n3 := &Node{2, node}
	n2 := &Node{3, []*Node{n5, n6}}
	root := &Node{1, []*Node{n2, n3, n4}}
	arr := LevelOrderII(root)
	fmt.Println(arr)
}