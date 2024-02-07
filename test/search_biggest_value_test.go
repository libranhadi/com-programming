package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBiggestValue(t *testing.T) {

	numbers := []int{1, -1, 5, 24, 10, 20, -20}
	result := 0

	for _, v := range numbers {
		if result > v {
			continue
		}
		result = v
	}

	assert.Equal(t, 24, result)
}
