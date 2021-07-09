package main

import (
	"fmt"
	"math"

	//"math"
)

//type List struct {
//	head *List//指向第一个节点
//}

type ListNode struct {
	Val int
	Next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var r1 int
//	key := 0
//	current := l1
//	for current.Next!=nil {
//		tmp := math.Pow10(key)
//		fmt.Println("%v", current.Val)
//		r1 += current.Val * int(tmp)
//		//fmt.Println("%d", r1)
//		current = current.Next
//		//fmt.Println("%v", current)
//		key++
//	}
//	//最后一个数字
//	key++
//	tmp := math.Pow10(key)
//	r1 = r1 + current.Val * int(tmp)
//	//fmt.Println("%d", r1)
//	//addTwoNumbers()
//
//	var r2 int
//	key2 := 0
//	current2 := l2
//	for current2.Next!=nil {
//		tmp := math.Pow10(key2)
//		r2 = r2 + current2.Val * int(tmp)
//		//fmt.Println("%d", current.Val)
//		current2 = current2.Next
//		//fmt.Println("%v", current)
//		key2++
//	}
//	//最后一个数字
//	key2++
//	tmp2 := math.Pow10(key2)
//	r2 = r2 + current2.Val * int(tmp2)
//	//fmt.Println("%d", r2)
//
//	result := r1+r2
//	//fmt.Println("%d", result)
//
//	//key3 := 100-1//最多100个数字,即最高位是10的99次方
//	//for i:=key3; i>=0; i-- {
//	//	tmp := math.Pow10(key3)
//	//	result/tmp
//	//}
//	result_s := strconv.Itoa(result)
//	//fmt.Println("%d", result)
//
//	result_map := strings.Split(result_s, "")
//	//fmt.Println("%v", result_map)
//
//	ResultListNode := &ListNode{}
//	//reverse_result_map := make(map[int]int)
//	mapleng := len(result_map)
//	var kv string
//	for i:=mapleng-1; i>=0; i-- {
//		kv = result_map[i]
//		kvi,_ := strconv.Atoi(kv)
//		if(i==mapleng-1) {
//			initList(ResultListNode, kvi)
//		} else {
//			addLinkNode(ResultListNode, kvi)
//		}
//		//fmt.Println("%v", kv)
//	}
//
//	return ResultListNode
//}
//var ResultListNode ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//fmt.Println("%v", l1.Next.Next)

	ResultListNode := &ListNode{}
	rln := ResultListNode
	//current := &ListNode{}
	//current.Next = l1
	//current2 := &ListNode{}
	//current2.Next = l2
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
		//fmt.Println("%v", current.Val)
		//fmt.Println("%v", current2.Val)
		//需要进一位
		if (result >= 10) {
			////还需要进位,但没有位置
			//if (current.Next == nil) {
			//	newNode := &ListNode{}
			//	newNode.Val = 1
			//	newNode.Next = nil
			//	current.Next = newNode
			//}
			//current.Next.Val += 1
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
	//return current
}

//if(current.Next==nil && current2.Next==nil) {
//fmt.Println("%v", current)
//back := &ListNode{}

//m := make(map[int]int)
//i := 0
//l_current := ResultListNode.Next
//for l_current.Next!=nil {
//	m[i] = l_current.Val
//	l_current = l_current.Next
//	i++
//}
//m[i] = l_current.Val
//fmt.Println("%v", m)
//length := len(m)
//b := &ListNode{}
////var b ListNode
////ll_current := b
////n 8 n.nex=nil b.nex=8
////n 0 n.nex=nil b.nex=0
////n 7 n.nex=nil b.nex=7
////for j:=0; j<=length-1; j++{
//for j:=length-1; j>=0; j--{
//
//	fmt.Println("%v", b)
//	fmt.Println("%v", m[j])
//	if (j == length-1) {
//		initList(b, m[j])
//	} else {
//		addLinkNode(b, m[j])
//		addLinkNode(b, m[j])
//	}
//
//	//if(j==length-1) {
//	//	initList(b, m[j])
//	//} else {
//	//	addLinkNode(b, m[j])
//	//}
//
//	//for ll_current.Next != nil {
//	//	ll_current = ll_current.Next
//	//}
//	//newNode := &ListNode{}
//	//newNode.Val = m[j]
//	//newNode.Next = nil
//	//ll_current.Next = newNode
//	//newNode.Next = ResultListNode.Next
//	//ResultListNode.Next = newNode
//}
//} else {
//	return addTwoNumbers(current.Next, current2.Next)
//}

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

	//if(l.Next == nil) {
	//	l.Next = newNode
	//} else {
		current.Next = newNode
	//}
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

	//addTwoNumbers()
}