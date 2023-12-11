// maxPointsOnLine.go

package main

import (
	"fmt"
	"math"
)

func main() {
	//读入
	nums := []int{2, 3, 4, 2, 5, 10}
	m := 4

	// 每个位置最大值，后续min比较
	dp := make([]int, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = math.MaxInt32
	}

	// 组合
	for i := 0; i < len(nums); i++ {
		for j := m; j >= nums[i]; j-- {
			// 倒叙遍历，保证只取一次
			dp[j] = min(dp[j], dp[j-nums[i]]+1)
		}
	}
	if dp[m] == math.MaxInt32 {
		fmt.Println("no solution")
	} else {
		fmt.Println(dp[m])
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
