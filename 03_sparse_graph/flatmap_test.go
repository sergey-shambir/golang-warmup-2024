package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMapWithIntKey(t *testing.T) {
	var m FlatMap[int, string]
	m.Add(5, "pear")
	m.Add(10, "banana")
	m.Add(7, "apple")

	assert.False(t, m.Contains(1))
	assert.True(t, m.Contains(5))
	assert.True(t, m.Contains(7))
	assert.True(t, m.Contains(10))

	assert.Equal(t, m.Get(1), "")
	assert.Equal(t, m.Get(5), "pear")
	assert.Equal(t, m.Get(7), "apple")
	assert.Equal(t, m.Get(10), "banana")
	assert.Equal(t, m.Get(8), "")
}

func TestFlatMapWithStringKey(t *testing.T) {
	var m FlatMap[string, string]
	m.Add("banana", "fruit")
	m.Add("apple", "unknown")
	m.Add("potato", "vegetable")
	m.Add("dog", "animal")

	// Add same value again
	m.Add("apple", "vegetable")
	m.Add("apple", "fruit")

	assert.Equal(t, m.Size(), 4)
	assert.Equal(t, m.Get("banana"), "fruit")
	assert.Equal(t, m.Get("apple"), "fruit")
	assert.Equal(t, m.Get("tomato"), "")
	assert.Equal(t, m.Get("potato"), "vegetable")
	assert.Equal(t, m.Get("dog"), "animal")
	assert.Equal(t, m.Get("cat"), "")

	m.Delete("dog")
	m.Delete("banana")
	assert.Equal(t, m.Size(), 2)

	// Preserved old data
	assert.Equal(t, m.Get("apple"), "fruit")
	assert.Equal(t, m.Get("tomato"), "")
	assert.Equal(t, m.Get("potato"), "vegetable")
	// Deleted two keys
	assert.Equal(t, m.Get("dog"), "")
	assert.Equal(t, m.Get("banana"), "")

	m.Add("cat", "animal")
	m.Add("turtle", "animal")
	m.Add("snail", "animal")
	m.Add("hamster", "animal")

	assert.Equal(t, m.Size(), 6)
	// Preserved old data
	assert.Equal(t, m.Get("apple"), "fruit")
	assert.Equal(t, m.Get("tomato"), "")
	assert.Equal(t, m.Get("potato"), "vegetable")
	// Deleted two keys
	assert.Equal(t, m.Get("dog"), "")
	assert.Equal(t, m.Get("banana"), "")
	// Added new data
	assert.Equal(t, m.Get("cat"), "animal")
	assert.Equal(t, m.Get("turtle"), "animal")
	assert.Equal(t, m.Get("snail"), "animal")
	assert.Equal(t, m.Get("hamster"), "animal")

	m.Add("tomato", "vegetable")
	assert.Equal(t, m.Size(), 7)
	assert.Equal(t, m.Get("apple"), "fruit")
	assert.Equal(t, m.Get("tomato"), "vegetable")
}
