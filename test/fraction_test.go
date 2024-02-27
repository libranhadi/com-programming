package test

import (
	"fmt"
	"testing"
)

type Denomination struct {
	Fraction int
	Qty      int
}

/**
 * @Author:
 * @Date: 2024-02-07 22:11:29
 * @Desc: Test Fraction
 */
func TestFraction(t *testing.T) {
	results := Fraction(2102)
	fmt.Println(results)
}

/**
 * @Author:
 * @Date: 2024-02-07 22:13:01
 * @Desc: Calculate fraction
 * @param INT eg : $2751
 * @fraction $1000, $100, $50, $20, $10, $5, $2, $1
 */
func Fraction(val int) []Denomination {
	denominations := make([]Denomination, 0)
	fractions := [8]int{1000, 100, 50, 20, 10, 5, 2, 1}

	for val > 0 {
		for _, v := range fractions {
			if val <= 0 {
				break
			}
			if val < v {
				continue
			}
			qty := val / v
			val %= v
			denominations = append(denominations, Denomination{Fraction: v, Qty: qty})
			break
		}
	}
	return denominations
}
