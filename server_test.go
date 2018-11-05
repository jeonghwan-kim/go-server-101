package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		x int
		y int
		z int
	}{
		{x: 1, y: 2, z: 3},
		{x: -1, y: -2, z: -3},
		{x: 1, y: -1, z: 0},
	}

	for _, test := range tests {
		actual := Sum(test.x, test.y)
		assert.Equal(t, actual, test.z, "they should be equal")
	}
}
