package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	m := 1
	n := 1
	for i := range obstacleGrid {
		dp[i] = make([]int, len(obstacleGrid[i]))

		if(obstacleGrid[i][0]==1 || m==0) {
			dp[i][0] = 0
			m = 0
		} else {
			dp[i][0] = 1
		}

		for j:=0; j<len(obstacleGrid[i]) ; j++ {
			if(obstacleGrid[0][j]==1 || n==0) {
				dp[0][j] = 0
				n = 0
			} else {
				dp[0][j] = 1
			}
		}
	}

	for i:=1; i<len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if(obstacleGrid[i][j]==1) {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func main()  {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	//obstacleGrid := [][]int{
	//	{0, 1},
	//	{0, 0},
	//}
	//obstacleGrid := [][]int{
	//	{1, 0},
	//}
	w := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(w)
}
