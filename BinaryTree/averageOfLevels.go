package main

import "fmt"

type TreeNodeAverageOfLevels struct {
	Val int
	Left *TreeNodeAverageOfLevels
	Right *TreeNodeAverageOfLevels
}

func averageOfLevels(root *TreeNodeAverageOfLevels) []float64 {
	var result []float64

	if root == nil {
		return result
	}

	var queue []*TreeNodeAverageOfLevels
	queue = append(queue, root)

	for len(queue)>0 {
		size := len(queue)

		var sum float64
		for i:=0; i<size; i++ {
			root = queue[0]
			queue = queue[1:]
			sum += float64(root.Val)
			if root.Left!=nil {
				queue = append(queue, root.Left)
			}
			if root.Right!=nil {
				queue = append(queue, root.Right)
			}
		}

		result = append(result, sum/float64(size))
	}

	return result
}

func main()  {
	tn5 := &TreeNodeAverageOfLevels{7, nil, nil}
	tn4 := &TreeNodeAverageOfLevels{15, nil, nil}
	tn3 := &TreeNodeAverageOfLevels{20, tn4, tn5}
	tn2 := &TreeNodeAverageOfLevels{9, nil, nil}
	root := &TreeNodeAverageOfLevels{3, tn2, tn3}
	arr := averageOfLevels(root)
	fmt.Println(arr)
}

