//https://leetcode.cn/problems/ones-and-zeroes/

package main

import (
	"fmt"
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i:=0; i<=m; i++ {
		dp[i] = make([]int, n+1)
		for j:=0; j<=n; j++ {
			dp[i][j] = 0
		}
	}

	for _, char := range strs {//遍历物品
		//zeroNum, oneNum := 0
		//计算字符中0跟1的数量
		zeroNum := strings.Count(char, "0")
		oneNumes := len(char) - zeroNum

		for i:=m; i>=zeroNum; i-- {
			for j:=n; j>=oneNumes; j-- {
				dp[i][j] = fmax(dp[i][j], dp[i-zeroNum][j-oneNumes]+1)
			}
		}
	}

	return dp[m][n]
}

func fmax(a int, b int) int {
	if a>b {
		return a
	}

	return b
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	v := findMaxForm(strs, m, n)
	fmt.Println(v)
}
