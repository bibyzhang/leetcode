package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//type List struct {
//	head *ListNode//指向第一个节点
//	size int
//}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r1 int
	key := 0
	current := l1
	for current.Next!=nil {
		tmp := math.Pow10(key)
		fmt.Println("%v", current.Val)
		r1 += current.Val * int(tmp)
		//fmt.Println("%d", r1)
		current = current.Next
		//fmt.Println("%v", current)
		key++
	}
	//最后一个数字
	key++
	tmp := math.Pow10(key)
	r1 = r1 + current.Val * int(tmp)
	//fmt.Println("%d", r1)
	//addTwoNumbers()

	var r2 int
	key2 := 0
	current2 := l2
	for current2.Next!=nil {
		tmp := math.Pow10(key2)
		r2 = r2 + current2.Val * int(tmp)
		//fmt.Println("%d", current.Val)
		current2 = current2.Next
		//fmt.Println("%v", current)
		key2++
	}
	//最后一个数字
	key2++
	tmp2 := math.Pow10(key2)
	r2 = r2 + current2.Val * int(tmp2)
	//fmt.Println("%d", r2)

	result := r1+r2
	//fmt.Println("%d", result)

	//key3 := 100-1//最多100个数字,即最高位是10的99次方
	//for i:=key3; i>=0; i-- {
	//	tmp := math.Pow10(key3)
	//	result/tmp
	//}
	result_s := strconv.Itoa(result)
	//fmt.Println("%d", result)

	result_map := strings.Split(result_s, "")
	//fmt.Println("%v", result_map)

	ResultListNode := &ListNode{}
	//reverse_result_map := make(map[int]int)
	mapleng := len(result_map)
	var kv string
	for i:=mapleng-1; i>=0; i-- {
		kv = result_map[i]
		kvi,_ := strconv.Atoi(kv)
		if(i==mapleng-1) {
			initList(ResultListNode, kvi)
		} else {
			addLinkNode(ResultListNode, kvi)
		}
		//fmt.Println("%v", kv)
	}

	return ResultListNode
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

	if(l.Next == nil) {
		l.Next = newNode
	} else {
		current.Next = newNode
	}
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

	initList(&ListNode1, 1)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 0)
	addLinkNode(&ListNode1, 1)


	var ListNode2 ListNode
	initList(&ListNode2, 9)
	////List2.addLinkNode(5)
	////List2.addLinkNode(6)
	////List2.addLinkNode(4)
	//
	////List2.addLinkNode(0)
	//
	addLinkNode(&ListNode2, 9)
	addLinkNode(&ListNode2, 9)
	addLinkNode(&ListNode2,9)

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

	//fmt.Println("%v", ListNode1)
	//fmt.Println("%v", ListNode2)
	//
	//for ListNode ListNode := range List1 {
	//
	//}

	result := addTwoNumbers(&ListNode1, &ListNode1)
	fmt.Println("%v", result.Next.Next.Next)

	//addTwoNumbers()
}