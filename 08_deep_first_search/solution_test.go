package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAchievePoint(t *testing.T) {
	cells := [][]int{
		{6, 4, 6, 3, 4, 7, 7, 5, 3, 3},
		{9, 3, 4, 3, 6, 7, 5, 8, 2, 5},
		{3, 2, 7, 3, 7, 3, 7, 2, 6, 4},
		{6, 7, 4, 7, 6, 7, 8, 8, 2, 2},
		{3, 3, 2, 8, 6, 7, 3, 3, 1, 8},
		{7, 4, 7, 4, 4, 5, 8, 6, 1, 6},
		{4, 7, 6, 8, 4, 7, 7, 7, 2, 3},
		{3, 1, 5, 3, 4, 2, 7, 2, 5, 6},
		{2, 4, 4, 9, 8, 3, 8, 6, 7, 8},
		{1, 2, 4, 3, 5, 1, 5, 9, 6, 6},
	}

	assert.True(t, CanAchievePoint(cells, Point{0, 0}, Point{6, 0}))
	assert.True(t, CanAchievePoint(cells, Point{0, 0}, Point{6, 7}))
	assert.True(t, CanAchievePoint(cells, Point{0, 0}, Point{0, 7}))
	assert.False(t, CanAchievePoint(cells, Point{1, 1}, Point{10, 6}))
}
