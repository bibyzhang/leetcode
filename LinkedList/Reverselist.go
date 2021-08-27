package main

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if(head.Next==nil) {
		return head
	}
	point := head
	var listMap = make(map[int]*ListNode)
	i := 0
	for nil!=point {
		if(i==0) {
			listMap[i] = head
		} else {
			listMap[i] = listMap[i-1].Next
		}
		point = point.Next
		i++
	}

	//fmt.Println(listMap)

	var lin = &ListNode{}
	k := len(listMap)-1
	//for j:=len(listMap)-1; j>=0; j-- {
	var ln_add = make(map[int]*ListNode)
	for j:=0; j<len(listMap); j++ {
		//fmt.Println(listMap[j-1])
		if(j==0) {
			lin = &ListNode{
				Val: listMap[j].Val,
				Next: nil,
			}
			ln_add[j] = lin
		} else {
			lin = &ListNode {
				Val: listMap[j].Val,
				Next: ln_add[j-1],
			}
			ln_add[j] = lin
		}
		k--
	}

	return lin
}

func main()  {
	var ln = &ListNode{}
	//linkArr := [5]int{1,2,3,4,5}
	linkArr := [2]int{1,2}
	var ln_add = make(map[int]*ListNode)

	for i := len(linkArr)-1; i >= 0 ; i-- {
		//最后一个节点
		if(i==len(linkArr)-1) {
			ln = &ListNode {
				Val: linkArr[i],
				Next: nil,
			}
			ln_add[i] = ln
		} else {
			//fmt.Println(ln_add[i+1])
			ln = &ListNode {
				Val: linkArr[i],
				Next: ln_add[i+1],
			}
			ln_add[i] = ln
		}
	}

	//fmt.Println(ln.Val)

	reverseList(ln)

	//fmt.Println(ln.Val)
}
