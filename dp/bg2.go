package main

import "fmt"

//有一个背包，能放最大重量为4
//有三个物品,其重量和价值如下
//A: 1[重量] 15[价值]
//A: 3[重量] 20[价值]
//A: 4[重量] 30[价值]
//求该背包能存放商品的最大价值
func maxValueBg2() int {
	maxBagSize := 4
	weight := []int{1, 3, 4}
	wv := []int{15, 20, 30}
	dp := make([]int, maxBagSize+1)

	//初始化
	for j:=0; j<=maxBagSize; j++ {
		dp[j] = 0
	}

	for i:=0; i<len(weight); i++ {
		for j:=maxBagSize; j>=weight[i]; j-- {
			dp[j] = bg2max(dp[j], dp[j-weight[i]] + wv[i])
		}
	}

	fmt.Println(dp)

	return dp[maxBagSize]
}

func bg2max(a int, b int) int {
	if a>b {
		return a
	}

	return b
}

func main()  {
	mv := maxValueBg2()
	fmt.Println(mv)
}
