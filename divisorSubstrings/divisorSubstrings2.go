//https://leetcode.cn/problems/find-the-k-beauty-of-a-number/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func divisorSubstrings(num int, k int) int {
	str := strconv.Itoa(num)
	s := len(str)
	//n := make([]int, s)
	c := 0
	//for _k, ch1 := range str {
	//	n[_k],_ = fmt.Println(string(ch1))
	//}
	str_arr := strings.Split(str, "")
	for i:=0; i<=s-k; i++ {
		f := str_arr[i]
		l := ""
		for j:=i+1; j<=i+k-1; j++ {
			l = l + str_arr[j]
		}

		e := f + l
		ei, _ := strconv.Atoi(e)
		fmt.Println(ei)
		if ei>0 && num%ei==0 {
			c++
		}
	}
	return c
}

func main()  {
	num := 430043
	k := 2
	v := divisorSubstrings(num, k)
	fmt.Println(v)
}