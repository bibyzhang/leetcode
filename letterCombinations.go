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
	//fmt.Println("%v", setp)

	//循环字母做顶点
	//for _,letter := range letters_arr {
	//	next_letter := dfsLetter(letters_arr, setp)
	//	data = append(data, next_letter)
	//}
	data = dfsLetter(letters_arr, setp)
	fmt.Println("v%", data)

	//match_letter_arr := strings.Split(match_letter, "")
	//count := 0
	//_count := 0
	//_setp_c := 0
	//_setp := 0
	//for _,letter := range match_letter_arr {
	//	if(count==0) {
	//		_setp = setp[0] //跳n步匹配
	//	}
	//	if(count==_setp) {
	//		_setp_c++
	//		_setp = setp[_setp_c]
	//		count = 0//重置
	//		_count+=_setp//没走完一个数字[步长]再加一个步长
	//	}
	//
	//	if(len(digits_arr)==1) {
	//		data = append(data, letter)
	//	} else {
	//		for i := _count + _setp; i < len(match_letter_arr); i++ {
	//			match_next_letter := match_letter_arr[i]
	//			r := letter + match_next_letter
	//			data = append(data, r)
	//		}
	//	}
	//	count++
	//}

	return data
}

func dfsLetter(letters_arr []string, setp map[string]int) []string {
	var _step int
	//fmt.Println("v%", setp)
	var data []string
	//var res []string
	//var _match_letter_arr []string
	//var r string
	//key := 0
	t := 0
	//panduan := 0
	//count_step := 0
	for __key,letter := range letters_arr {
		_step = setp[letter]
		//判断够不够截取

		//顶点只循环第一轮[即第一个数字]
		//超过一个数字的放到递归中
		if(__key==0) {
			t = len(letters_arr[:_step])
			//count_step = _step
		}
		if(t!=0 && __key>=t) {
			continue
		}
		//if(__key>=_step) {
		//	continue
		//}

		//res = nil
		//if(__key<_step) {
		//if(_step==3){
		//	if((__key+_step)<len(letters_arr)) {
		//		//去除步长之后的字母继续循环
		//		_match_letter_arr = letters_arr[_step:]
		//		res = dfsLetter(_match_letter_arr, setp)
		//		fmt.Println("v%", _match_letter_arr)
		//	} else if(__key<_step) {
		//		data = append(data, letter)
		//	}
		//} else if(_step==4) {
		//	if((__key+_step)<=len(letters_arr)) {
		//		//去除步长之后的字母继续循环
		//		_match_letter_arr = letters_arr[_step:]
		//		res = dfsLetter(_match_letter_arr, setp)
		//		fmt.Println("v%", _match_letter_arr)
		//	} else if(__key<_step) {
		//		data = append(data, letter)
		//	}
		//}

		//要不要截取依赖下一个长度 非当前长度
		//if(__key+_step<len(letters_arr)) {
		//fmt.Println("v%", _step)
		//fmt.Println("v%", len(letters_arr[_step:]))
		//存在下一个数
		if(len(letters_arr[_step:])>0) {
		//if((_step+count_step)<len(letters_arr)) {
			//去除步长之后的字母继续循环
			_match_letter_arr := letters_arr[_step:]
			res := dfsLetter(_match_letter_arr, setp)
			//fmt.Println("v%", _step)
			//fmt.Println("v%", _step)
			//fmt.Println("v%", _match_letter_arr)

			for _, _key := range res {
				r := letter + _key
				data = append(data, r)
			}
		//} else if(__key<=_step) {
		} else {
			data = append(data, letter)
			//没有产生递归返回数据,即当前只有一个数字
			//if(res==nil) {
			//	data = append(data, letter)
			//} else {//递归返回部分,超过一个数字
			//	for _, _key := range res {
			//		r = letter + _key
			//		data = append(data, r)
			//	}
			//}
		}
		//count_step++


		//到最小单位才组合,然后返回给上一层
		//if(len(res)==0) {
		//	data = append(data, letter)
		//} else {
		//	if(res!=nil) {
		//		for _, _key := range res {
		//			r = letter + _key
		//			data = append(data, r)
		//		}
		//	}
		//}
		//if(__key<=_step && res!=nil) {
		//	for _, _key := range res {
		//		r = letter + _key
		//		data = append(data, r)
		//	}
		//}
		//key++

		//else {
		//	return _match_letter_arr
		//}

		//不跟步长内的字母组合
		//if(key>=_step) {
		//	s += letter
		//	data = append(data, s)
		//}


		//步长之后没有值了,进入循环拼接
		//fmt.Println("v%", letters_arr[key+_step])
		//r += letter
		//data = append(data, r)
		//fmt.Println("v%", r)
	}

	return data
}

func main() {
	digits := "6789"
	letterCombinations(digits)
	//fmt.Println("v%", res)
}
