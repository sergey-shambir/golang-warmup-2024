package main

import "slices"

type orderedKey interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~string
}

type keyValue[K orderedKey, V any] struct {
	key   K
	value V
}

type FlatMap[K orderedKey, V any] struct {
	items []keyValue[K, V]
}

func (m FlatMap[K, V]) Items() []keyValue[K, V] {
	return m.items
}

func (m *FlatMap[K, V]) Size() int {
	return len(m.items)
}

func (m *FlatMap[K, V]) Get(key K) V {
	var result V
	index := m.findPartitionPoint(key)
	if index != len(m.items) && m.items[index].key == key {
		result = m.items[index].value
	}
	return result
}

func (m *FlatMap[K, V]) Contains(key K) bool {

	index := m.findPartitionPoint(key)
	return index < len(m.items) && m.items[index].key == key
}

func (m *FlatMap[K, V]) Add(key K, value V) {
	index := m.findPartitionPoint(key)
	if index != len(m.items) && m.items[index].key == key {
		m.items[index].value = value
	} else {
		m.items = slices.Insert(m.items, index, keyValue[K, V]{key, value})
	}
}

// Returns least index where (m.items[index].key >= key)
func (m *FlatMap[K, V]) findPartitionPoint(key K) int {
	left, right := 0, len(m.items)
	for left != right {
		middle := (left + right) / 2
		if m.items[middle].key < key {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}

func (m *FlatMap[K, V]) Delete(key K) {
	index := m.findPartitionPoint(key)
	if index != len(m.items) && m.items[index].key == key {
		m.items = append(m.items[0:index], m.items[index+1:]...)
	}
}
