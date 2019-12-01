package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFib(t *testing.T) {
	assert.Equal(t, int64(0), fib(0))
	assert.Equal(t, int64(1), fib(1))
	assert.Equal(t, int64(1), fib(2))
	assert.Equal(t, int64(6765), fib(20))
}

func TestFibAll(t *testing.T) {
	for i := int64(15); i < 40; i++ {
		fibNumbers := fibAll(i)
		require.NotNil(t, fibNumbers)
		require.Len(t, fibNumbers, int(i+1))
		assert.Equal(t, 233, fibNumbers[13])
		assert.Equal(t, 55, fibNumbers[10])
	}
}
