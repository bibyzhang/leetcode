//https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

package main

import "fmt"

type LinkNode2 struct {
	Val int
	Next *LinkNode2
}

func reverseList2(head *LinkNode2) *LinkNode2 {
	//空节点,直接返回
	//单个节点,直接返回
	if head == nil || head.Next == nil {
		return head
	}

	c := head.Next
	//头节点
	prev := head
	prev.Next = nil
	for c != nil {
		n := head.Next
		n.Next = prev
		prev = c
		c = n
	}

	return prev
}

func main()  {
	var ln = &LinkNode2{}
	//linkArr := [5]int{1,2,3,4,5}
	linkArr := [2]int{1,2}
	var ln_add = make(map[int]*LinkNode2)

	for i := len(linkArr)-1; i >= 0 ; i-- {
		//最后一个节点
		if(i==len(linkArr)-1) {
			ln = &LinkNode2 {
				Val: linkArr[i],
				Next: nil,
			}
			ln_add[i] = ln
		} else {
			//fmt.Println(ln_add[i+1])
			ln = &LinkNode2 {
				Val: linkArr[i],
				Next: ln_add[i+1],
			}
			ln_add[i] = ln
		}
	}
	reverseList2(ln)
	fmt.Println(ln.Val)
}