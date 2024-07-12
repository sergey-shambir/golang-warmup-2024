package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAchievableCellsWithTrivialCases(t *testing.T) {
	assert.Equal(t, countAchievableCells(point{0, 0}, 0), 1)
	assert.Equal(t, countAchievableCells(point{0, 0}, 1), 5)
	assert.Equal(t, countAchievableCells(point{0, 1}, 1), 5)
	assert.Equal(t, countAchievableCells(point{1, 1}, 2), 13)
	assert.Equal(t, countAchievableCells(point{0, 0}, 2), 13)
}

func TestCountAchievableCellsWithMaxDigitsSum25(t *testing.T) {
	assert.Equal(t, countAchievableCells(point{1000, 1000}, 25), 148848)
}
