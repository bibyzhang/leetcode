//https://leetcode.cn/problems/unique-paths/submissions/

package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i:=0; i<m; i++ {
		dp[i][0] = 1
	}
	for j:=0; j<n; j++ {
		dp[0][j] = 1
	}

	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main()  {
	p := uniquePaths(1, 1)
	fmt.Println(p)
}