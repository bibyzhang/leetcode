package main

//import "fmt"
import (
	"./LinkedList"
	"fmt"
)

func main()  {
	var ln = &LinkedList.ListNode{}
	linkArr := [5]int{1,2,3,4,5}
	//linkArr := [2]int{1,2}
	var ln_add = make(map[int]*LinkedList.ListNode)

	for i := len(linkArr)-1; i >= 0 ; i-- {
		//最后一个节点
		if(i==len(linkArr)-1) {
			ln = &LinkedList.ListNode {
				Val: linkArr[i],
				Next: nil,
			}
			ln_add[i] = ln
		} else {
			//fmt.Println(ln_add[i+1])
			ln = &LinkedList.ListNode {
				Val: linkArr[i],
				Next: ln_add[i+1],
			}
			ln_add[i] = ln
		}
	}

	LinkedList.Reverse(ln)

	fmt.Println(ln.Next)
}
