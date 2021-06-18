//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

//第一种错误思路
//func letterCombinations(digits string) []string {
//	var data []string
//	digits_arr := strings.Split(digits, "")
//	start := 97//a的ASCII十进制值
//	ca_len := 3//每个数字对应的字母长度为3,z特殊对应四位
//	fortimes := len(digits_arr)-1
//
//	lower := 0
//	for _,val := range digits_arr {
//		var_int,err:=strconv.Atoi(val)
//		if(err!=nil) {
//			return data
//		}
//		//7、9对应4个字母
//		if(var_int==9 || var_int==7) {
//			ca_len = ca_len+1
//		}
//
//		start_distrbute_num := (var_int - 2)*3
//		if(var_int==8 || var_int==9) {
//			start_distrbute_num = start_distrbute_num+1
//		}
//		rel_start := start + start_distrbute_num
//		for i := 0; i < ca_len; i++ {
//			match_letter_num := rel_start + i
//			match_letter := string(rune(match_letter_num))
//
//			if(len(digits_arr)==1) {
//				data = append(data, match_letter)
//			} else {
//				for j := 0; j < (ca_len * (fortimes)); j++ {
//					match_next_letter := string(rune(rel_start + ca_len + j))
//					r := match_letter + match_next_letter
//					data = append(data, r)
//					lower++
//				}
//			}
//		}
//		fortimes--
//	}
//
//	return data
//}

//第二种错误思路
//func letterCombinations(digits string) []string {
//	var data []string
//	digits_arr := strings.Split(digits, "")
//
//	start := 97//a的ASCII十进制值
//	ca_len := 3//每个数字对应的字母长度为3,z特殊对应四位
//	var match_letter string
//	var setp []int
//
//	for _,val := range digits_arr {
//		var_int,err := strconv.Atoi(val)
//
//		if(err!=nil) {
//			return data
//		}
//		//7、9对应4个字母
//		if(var_int==9 || var_int==7) {
//			ca_len = ca_len+1
//		}
//		setp = append(setp, ca_len)
//
//		start_distrbute_num := (var_int - 2)*3
//		if(var_int==8 || var_int==9) {
//			start_distrbute_num = start_distrbute_num+1
//		}
//		rel_start := start + start_distrbute_num
//		for i := 0; i < ca_len; i++ {
//			match_letter_num := rel_start + i
//			match_letter += string(rune(match_letter_num))
//		}
//	}
//
//	match_letter_arr := strings.Split(match_letter, "")
//	count := 0
//	_count := 0
//	_setp_c := 0
//	_setp := 0
//	for _,letter := range match_letter_arr {
//		if(count==0) {
//			_setp = setp[0] //跳n步匹配
//		}
//		if(count==_setp) {
//			_setp_c++
//			_setp = setp[_setp_c]
//			count = 0//重置
//			_count+=_setp//没走完一个数字[步长]再加一个步长
//		}
//
//		if(len(digits_arr)==1) {
//			data = append(data, letter)
//		} else {
//			for i := _count + _setp; i < len(match_letter_arr); i++ {
//				match_next_letter := match_letter_arr[i]
//				r := letter + match_next_letter
//				data = append(data, r)
//			}
//		}
//		count++
//	}
//
//	return data
//}

func letterCombinations(digits string) []string {
	var data []string
	digits_arr := strings.Split(digits, "")

	start := 97//a的ASCII十进制值
	var ca_len  int//每个数字对应的字母长度为3,z特殊对应四位
	var match_letter string
	setp := make(map[string]int)

	//拿到所有的字母
	for _,val := range digits_arr {
		var_int,err := strconv.Atoi(val)

		if(err!=nil) {
			return data
		}
		//7、9对应4个字母
		ca_len = 3
		if(var_int==9 || var_int==7) {
			ca_len = ca_len+1
		}

		start_distrbute_num := (var_int - 2)*3
		if(var_int==8 || var_int==9) {
			start_distrbute_num = start_distrbute_num+1
		}

		rel_start := start + start_distrbute_num
		for i := 0; i < ca_len; i++ {
			match_letter_num := rel_start + i
			letter := string(rune(match_letter_num))
			setp[letter] = ca_len
			match_letter += letter
		}
	}

	letters_arr := strings.Split(match_letter, "")
	data = dfsLetter(letters_arr, setp)
	return data
}

func dfsLetter(letters_arr []string, setp map[string]int) []string {
	var _step int
	var data []string
	t := 0
	for __key,letter := range letters_arr {
		_step = setp[letter]

		//顶点只循环第一轮[即第一个数字]
		//超过一个数字的放到递归中
		if(__key==0) {
			t = len(letters_arr[:_step])
			//count_step = _step
		}
		if(t!=0 && __key>=t) {
			continue
		}

		//存在下一个数
		if(len(letters_arr[_step:])>0) {
			//去除步长之后的字母继续循环
			_match_letter_arr := letters_arr[_step:]
			res := dfsLetter(_match_letter_arr, setp)
			for _, _key := range res {
				r := letter + _key
				data = append(data, r)
			}
		} else {
			//没有产生递归返回数据,即当前只有一个数字
			data = append(data, letter)
		}
	}

	return data
}

func main() {
	digits := "6789"
	data := letterCombinations(digits)
	fmt.Println("v%", data)
}
