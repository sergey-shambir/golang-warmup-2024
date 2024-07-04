package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvlTreeMapWithIntKeys(t *testing.T) {
	var m AvlTreeMap[int, string]
	m.Insert(10, "banana")

	assert.Equal(t, m.Keys(), []int{10})
	assert.Equal(t, m.Values(), []string{"banana"})

	m.Insert(12, "coconut")
	m.Insert(3, "cherry")

	assert.Equal(t, m.Keys(), []int{3, 10, 12})
	assert.Equal(t, m.Values(), []string{"cherry", "banana", "coconut"})

	m.Insert(7, "blueberry")
	m.Insert(1, "lemon")
	m.Insert(4, "pineapple")
	m.Insert(15, "pear")

	assert.Equal(t, []int{1, 3, 4, 7, 10, 12, 15}, m.Keys())
	assert.Equal(t, []string{"lemon", "cherry", "pineapple", "blueberry", "banana", "coconut", "pear"}, m.Values())

	m.Remove(3)

	assert.Equal(t, []int{1, 4, 7, 10, 12, 15}, m.Keys())
	assert.Equal(t, []string{"lemon", "pineapple", "blueberry", "banana", "coconut", "pear"}, m.Values())

	m.Remove(10)
	m.Remove(7)
	m.Remove(12)

	assert.Equal(t, m.Keys(), []int{1, 4, 15})
	assert.Equal(t, m.Values(), []string{"lemon", "pineapple", "pear"})

	m.Insert(7, "raspberry")
	m.Insert(8, "strawberry")
	m.Remove(1)

	assert.Equal(t, m.Keys(), []int{4, 7, 8, 15})
	assert.Equal(t, m.Values(), []string{"pineapple", "raspberry", "strawberry", "pear"})
}
