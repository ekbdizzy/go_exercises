package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasSameValues(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 5}
	assert.True(t, hasSameValues(nums))
	nums = []int{1, 2, 3, 4, 5, 6}
	assert.False(t, hasSameValues(nums))
}

func TestIsSorted(t *testing.T) {
	words := []string{"a", "ab", "ac", "ad"}
	assert.True(t, isSorted(words))

	words = []string{"a", "c", "ac", "ad"}
	assert.False(t, isSorted(words))

	words = []string{"c", "daddy", "format", "hello"}
	assert.True(t, isSorted(words))

	words = []string{"zeppelin", "zoo", "zebra"}
	assert.False(t, isSorted(words))
}
