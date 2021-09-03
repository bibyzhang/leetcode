package LinkedList

import "fmt"

type LinkList struct {
	Head *LinkNode
}

type LinkNode struct {
	Val interface{}
	Next *LinkNode
}

func Reverse(head *ListNode) {
	//头部节点为nil 空链表
	//第一个节点为nil 空链表
	//只有一个节点
	//if(head == nil || head.Next==nil || head.Next.Next==nil) {

	//current := head.Next
	//prev := &ListNode{}
	//for current!=nil {
	//	tmpnode := current.Next
	//	current.Next = prev
	//	prev = current
	//	current = tmpnode
	//}
	//prev := &ListNode{}


	if(head == nil || head.Next==nil) {
		return
	}

	current := head.Next
	prev := head
	prev.Next = nil
	for current!=nil {
		tmpnode := current.Next
		current.Next = prev
		prev = current
		current = tmpnode
	}

	head = prev

	fmt.Println(head)


	return
}


