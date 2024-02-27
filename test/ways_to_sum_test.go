package test

import (
	"fmt"
	"testing"
)

func TestWaysToSum(t *testing.T) {
	total := int(5)
	k := int(3)

	ways := wayToSum(total, k)

	fmt.Println("Result:", ways)
}

func wayToSum(total, k int) int {
	dp := make([]int, total+1)
	dp[0] = 1

	for i := 1; i <= total; i++ {
		for j := 1; j <= k; j++ {
			if i >= j {
				dp[i] += dp[i-j]
			}
		}
	}

	return dp[total]
}
