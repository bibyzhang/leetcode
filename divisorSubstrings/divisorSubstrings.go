//https://leetcode.cn/problems/find-the-k-beauty-of-a-number/

package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings2(num int, k int) int {
	str := strconv.Itoa(num) //先转字符串
	s := len(str)
	c := 0
	for i:=0; i<=s-k; i++ {
		e, _ := strconv.Atoi(str[i:i+k])//截取字符串
		if e>0 && num%e==0 {
			c++
		}
	}
	return c
}

func main()  {
	num := 240
	k := 2
	v := divisorSubstrings2(num, k)
	fmt.Println(v)
}