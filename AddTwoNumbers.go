//https://leetcode-cn.com/problems/add-two-numbers/submissions/

package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ResultListNode := &ListNode{}
	rln := ResultListNode
	current := l1
	current2 := l2
	is_larage_than_10 := 0
	for current != nil || current2 != nil {
		//从个位开始加,然后往高位递归
		if (current.Next == nil && current2.Next != nil) {
			node := &ListNode{}
			node.Val = 0
			node.Next = nil
			current.Next = node
		}
		if (current.Next != nil && current2.Next == nil) {
			node := &ListNode{}
			node.Val = 0
			node.Next = nil
			current2.Next = node
		}

		result := current.Val + current2.Val + is_larage_than_10
		//需要进一位
		if (result >= 10) {
			is_larage_than_10 = 1
		} else {
			is_larage_than_10 = 0
		}

		val := int(math.Mod(float64(result), float64(10)))
		nNode := &ListNode{}
		nNode.Val = val
		nNode.Next = nil
		//找到最后一个节点[添加节点前]
		for rln.Next!=nil {
			rln = rln.Next
		}
		rln.Next = nNode

		current = current.Next
		current2 = current2.Next
	}

	if(is_larage_than_10==1) {
		nNode := &ListNode{}
		nNode.Val = 1
		nNode.Next = nil
		//找到最后一个节点[添加节点前]
		for rln.Next!=nil {
			rln = rln.Next
		}
		rln.Next = nNode
	}

	return ResultListNode.Next
}

func initList(l *ListNode, val int)  {
	l.Val = val
	l.Next = nil
}

func addLinkNode(l *ListNode, val int)  {
	newNode := &ListNode{}
	newNode.Val = val
	newNode.Next = nil

	//找到最后一个节点
	current := l
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func main() {
	var ListNode1 ListNode
	//initList(&ListNode1, 9)
	//List1.addLinkNode(2)
	//List1.addLinkNode(4)
	//List1.addLinkNode(3)

	//List1.addLinkNode(0)

	//addLinkNode(&ListNode1,9)
	//addLinkNode(&ListNode1,9)
	//addLinkNode(&ListNode1,9)
	//addLinkNode(&ListNode1,9)
	//addLinkNode(&ListNode1,9)
	//addLinkNode(&ListNode1,9)

	//测试1
	initList(&ListNode1, 2)
	addLinkNode(&ListNode1, 4)
	addLinkNode(&ListNode1, 3)

	//测试2
	//initList(&ListNode1, 1)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 0)
	//addLinkNode(&ListNode1, 1)

	//测试3
	//initList(&ListNode1, 0)

	var ListNode2 ListNode

	//测试1
	initList(&ListNode2, 5)
	addLinkNode(&ListNode2,6)
	addLinkNode(&ListNode2,4)
	//
	////测试2
	//initList(&ListNode2, 5)
	//addLinkNode(&ListNode2,6)
	//addLinkNode(&ListNode2,4)
	//addLinkNode(&ListNode2,4)
	//

	//测试3
	//addLinkNode(&ListNode2,0)


	////List2.addLinkNode(0)
	//
	//addLinkNode(&ListNode2, 6)
	//addLinkNode(&ListNode2, 4)
	//addLinkNode(&ListNode2,9)

	//var ListNode1 ListNode
	//ListNode1.Val = 2
	//ListNode1.Next = nil
	//
	//newNode := &ListNode{}
	//newNode.Val = 2
	//newNode.Next = nil
	//
	//ListNode1.addLinkNode(2)
	//ListNode1.addLinkNode(4)
	//ListNode1.addLinkNode(3)
	//
	//var ListNode2 ListNode
	//ListNode2.addLinkNode(5)
	//ListNode2.addLinkNode(6)
	//ListNode2.addLinkNode(7)

	//fmt.Println("%v", ListNode2.Next.Next)
	//fmt.Println("%v", ListNode2)
	//
	//for ListNode ListNode := range List1 {
	//
	//}

	result := addTwoNumbers(&ListNode1, &ListNode2)
	fmt.Println("%v", result.Next.Next)
}