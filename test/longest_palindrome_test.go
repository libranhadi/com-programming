package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func LongestPalindrome(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}

	matArray := make([][]bool, n)
	for i := range matArray {
		matArray[i] = make([]bool, n)
		matArray[i][i] = true
	}

	start, maxLength := 0, 1
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && (length == 2 || matArray[i+1][j-1]) {
				matArray[i][j] = true
				if length > maxLength {
					start = i
					maxLength = length
				}
			}
		}
	}
	return s[start : start+maxLength]
}

func TestLongestPalindrome(t *testing.T) {
	a := LongestPalindrome("babad")

	assert.Equal(t, "bab", a)
}
