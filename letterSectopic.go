//https://leetcode-cn.com/problems/valid-anagram/solution/

package main

func isAnagram(s string, t string) bool {
	//字母异位单词长度一致
	s_len := len(s)
	t_len := len(t)
	if(s_len != t_len) {
		return false
	}

	ms1 := make(map[byte]byte)
	ms2 := make(map[byte]byte)

	//计算单词中相同ASCII码的值相同
	for i := 0; i<len(s); i++ {
		s1 := s[i]
		var s_asciiBytes byte = byte(s1)
		ms1[s_asciiBytes] += s_asciiBytes
	}

	for i := 0; i<len(s); i++ {
		t1 := t[i]
		var t_asciiBytes byte = byte(t1)
		ms2[t_asciiBytes] += t_asciiBytes
	}

	for k, v := range ms1 {
		if(v!=ms2[k]) {
			return false
		}
	}

	return true
}

func main()  {
	s := "anagram"
	t := "nagaram"
	//s := "rat"
	//t := "cat"
	//s := "aacc"
	//t := "ccac"
	//s := "aacc"
	//t := "ccac"
	//s := "ac"
	//t := "bb"
	r := isAnagram(s, t)
	if(r) {
		println("true")
	} else {
		println("false")
	}
}