package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	equal := GCD(12, 12)
	assert.Equal(t, 12, equal)

	firstZero := GCD(0, 5)
	assert.Equal(t, 5, firstZero)

	secondZero := GCD(3, 0)
	assert.Equal(t, 3, secondZero)

	withOne := GCD(1, 17)
	assert.Equal(t, 1, withOne)

	regular := GCD(20, 30)
	assert.Equal(t, 10, regular)
	regularReverse := GCD(30, 20)
	assert.Equal(t, 10, regularReverse)
}