package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeParts(t *testing.T) {
	chunk := []int { 5, 8, 6, 7, 12 }
	MergeParts(chunk, 0, 2, 5)

	expected := []int { 5, 6, 7, 8, 12 }
	
	fmt.Println(chunk)
	assert.Equal(t, chunk, expected, "list must be sorted")
}