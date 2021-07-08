package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type List struct {
	head *ListNode//指向第一个节点
	size int
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//fmt.Println("%v", l1)

	var r1 int
	key := 0
	current := l1
	for current.Next!=nil {
		tmp := math.Pow10(key)
		r1 = r1 + current.Val * int(tmp)
		//fmt.Println("%d", current.Val)
		current = current.Next
		//fmt.Println("%v", current)
		key++
	}
	//最后一个数字
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

	ResultList := &List{}
	ResultList.initList()
	//reverse_result_map := make(map[int]int)
	mapleng := len(result_map)
	var kv string
	for i:=mapleng-1; i>=0; i-- {
		kv = result_map[i]
		kvi,_ := strconv.Atoi(kv)
		ResultList.addLinkNode(kvi)
		//fmt.Println("%v", kv)
	}
	//for key,val := range result_map {
	//
	//}

	return ResultList.head
}

func (l *List) initList()  {
	head := &ListNode{}
	l.head = head
	l.size = 0
}

func (l *List) addLinkNode(val int)  {
	current := l.head
	//找到最后一个节点
	for current.Next != nil {
		current = current.Next
	}
	newNode := &ListNode{}
	newNode.Val = val
	newNode.Next = nil
	current.Next = newNode
	l.size++
}

//func (l *ListNode) initList(val int)  {
//	head := &ListNode{}
//	l.Val = val
//	l.Next = nil
//}
//
//func (l *ListNode) addLinkNode(val int)  {
//	//找到最后一个节点
//	var current *ListNode
//	for l.Next != nil {
//		current = l.Next
//	}
//	newNode := &ListNode{}
//	newNode.Val = val
//	newNode.Next = nil
//	if(current!=nil) {//第一个元素不用
//		current.Next = newNode
//	}
//}

func main() {
	var List1 List
	List1.initList()
	//List1.addLinkNode(2)
	//List1.addLinkNode(4)
	//List1.addLinkNode(3)

	//List1.addLinkNode(0)

	List1.addLinkNode(9)
	List1.addLinkNode(9)
	List1.addLinkNode(9)
	List1.addLinkNode(9)
	List1.addLinkNode(9)
	List1.addLinkNode(9)
	List1.addLinkNode(9)

	var List2 List
	List2.initList()
	//List2.addLinkNode(5)
	//List2.addLinkNode(6)
	//List2.addLinkNode(4)

	//List2.addLinkNode(0)

	List2.addLinkNode(9)
	List2.addLinkNode(9)
	List2.addLinkNode(9)
	List2.addLinkNode(9)


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

	//fmt.Println("%v", List1)
	//fmt.Println("%v", List2)
	//
	//for ListNode ListNode := range List1 {
	//
	//}

	result := addTwoNumbers(List1.head.Next, List2.head.Next)
	fmt.Println("%v", result)

	//addTwoNumbers()
}