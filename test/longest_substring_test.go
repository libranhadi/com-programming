package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputs = []map[string]any{
	map[string]any{"input": "abcabcbb", "output": 3},
	map[string]any{"input": "bbbbbb", "output": 1},
	map[string]any{"input": "pwwwkew", "output": 3},
	map[string]any{"input": "kalsdamdsadd", "output": 5},
	map[string]any{"input": " ", "output": 1},
	map[string]any{"input": "  ", "output": 1},
}

/*
*

  - @Author: flydreame

  - @Date: 2024-02-11 12:07:22

  - @Desc: Technical Test Search Longest Substring

    Example 1:
    Input: s = "abcabcbb"
    Output: 3
    Explanation: The answer is "abc", with the length of 3.
    Example 2:

    Input: s = "bbbbb"
    Output: 1
    Explanation: The answer is "b", with the length of 1.
    Example 3:

    Input: s = "pwwkew"
    Output: 3
    Explanation: The answer is "wke", with the length of 3.
    Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

  - @Constraints:
    0 <= s.length <= 5 * 104
    s consists of English letters, digits, symbols and spaces.
*/
func TestSearchingLongestSubstring(t *testing.T) {
	for _, v := range inputs {
		// result := lengthOfLongestSubstring(v["input"].(string))
		// assert.Equal(t, v["output"], result)

		result2 := lengthOfLongestSubstringRefactoring(v["input"].(string))
		assert.Equal(t, v["output"], result2)
	}
}

func BenchmarkSearchingLongestSubstring(b *testing.B) {
	for _, v := range inputs {
		// result := lengthOfLongestSubstring(v["input"].(string))
		// fmt.Println(result)

		result2 := lengthOfLongestSubstringRefactoring(v["input"].(string))
		fmt.Println(result2)
	}
}

func lengthOfLongestSubstring(s string) int {
	var result string
	isContinue := false
	for i := 0; i < len(s); i++ {
		parentChar := s[i]
		subCharacters := s[i+1:]
		isContinue = false
		var tempResult string

		for j := 0; j < len(subCharacters); j++ {
			childChar := subCharacters[j]
			if parentChar == childChar || strings.ContainsAny(tempResult, string(childChar)) {
				isContinue = true
				break
			}
			tempResult += string(childChar)
		}

		tempResult = string(parentChar) + tempResult

		if len(result) < len(tempResult) {
			result = tempResult
		}

		if isContinue {
			continue
		}
	}
	return len(result)
}

func lengthOfLongestSubstringRefactoring(s string) int {
	LongestSubstr := 0
	WindowStart := 0
	seenRunes := map[int32]int{}

	for currentIndex, v := range s {
		if value, ok := seenRunes[v]; ok {
			if value > WindowStart-1 {
				WindowStart = max(seenRunes[v]+1, WindowStart+1)
			}
		}
		seenRunes[v] = currentIndex
		LongestSubstr = max(LongestSubstr, len(s[WindowStart:currentIndex+1]))
	}

	return LongestSubstr

}

// return highest value of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
