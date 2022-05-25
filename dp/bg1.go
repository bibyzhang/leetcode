package main

import "fmt"

//有一个背包，能放最大重量为4
//有三个物品,其重量和价值如下
//A: 1[重量] 15[价值]
//A: 3[重量] 20[价值]
//A: 4[重量] 30[价值]
//求该背包能存放商品的最大价值
func maxValueBg() int {
	maxBagWeight := 4
	weight := []int{1, 3, 4}
	bagValue := []int{15, 20, 30}
	dp := make([][]int, len(weight))
	for i:=0; i<len(weight); i++ {
		dp[i] = make([]int, maxBagWeight+1)
		for j := 0; j <= maxBagWeight; j++ {
			dp[i][j] = 0
		}
	}

	for j:=maxBagWeight; j>=weight[0]; j-- {
		dp[0][j] = dp[0][j - weight[0]] + bagValue[0]
	}

	for i:=1; i<len(weight); i++ {
		for j := 0; j <= maxBagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dpmax(dp[i-1][j], dp[i-1][j-weight[i]] + bagValue[i])
			}
		}
	}

	fmt.Println(dp)

	return dp[len(weight)-1][maxBagWeight]
}

func dpmax(a int, b int) int {
	if a>b {
		return a
	}
	return b
}

func main()  {
	max_value := maxValueBg()
	fmt.Println(max_value)
}