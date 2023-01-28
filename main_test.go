package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func funcTestHasSameValues(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 5}
	assert.Equal(t, hasSameValues(nums), false)
	nums = []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, hasSameValues(nums), true)
}
