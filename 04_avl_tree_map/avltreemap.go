package main

type AvlTreeMap[K orderedKey, V any] struct {
	root *avlTreeNode[K, V]
}

type orderedKey interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~string
}

type avlTreeNode[K orderedKey, V any] struct {
	key    K
	value  V
	height int8
	left   *avlTreeNode[K, V]
	right  *avlTreeNode[K, V]
}

func (m *AvlTreeMap[K, V]) Insert(key K, value V) {
	m.root = insertAvlTreeNode(m.root, key, value)
}

func (m *AvlTreeMap[K, V]) Remove(key K) {
	m.root = removeAvlTreeNode(m.root, key)
}

func (m *AvlTreeMap[K, V]) Keys() []K {
	return listAvlTreeKeys(m.root, nil)
}

func (m *AvlTreeMap[K, V]) Values() []V {
	return listAvlTreeValues(m.root, nil)
}

func (n *avlTreeNode[K, V]) balance() *avlTreeNode[K, V] {
	n.updateHeight()
	factor := n.balanceFactor()
	if factor == 2 {
		if n.left.balanceFactor() > 0 {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	if factor == -2 {
		if n.right.balanceFactor() < 0 {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	}
	return n
}

func (n *avlTreeNode[K, V]) findLeftmostChild() *avlTreeNode[K, V] {
	if n.left != nil {
		return n.left.findLeftmostChild()
	}
	return n
}

func (n *avlTreeNode[K, V]) removeLeftmostChild() *avlTreeNode[K, V] {
	if n.left == nil {
		return n.right
	}
	n.left = n.left.removeLeftmostChild()
	return n.balance()
}

func (n *avlTreeNode[K, V]) rotateRight() *avlTreeNode[K, V] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

func (n *avlTreeNode[K, V]) rotateLeft() *avlTreeNode[K, V] {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

func (n *avlTreeNode[K, V]) balanceFactor() int8 {
	return n.left.heightOrZero() - n.right.heightOrZero()
}

func (n *avlTreeNode[K, V]) updateHeight() {
	n.height = 1 + max(n.left.heightOrZero(), n.right.heightOrZero())
}

func (n *avlTreeNode[K, V]) heightOrZero() int8 {
	if n != nil {
		return n.height
	}
	return 0
}

func insertAvlTreeNode[K orderedKey, V any](n *avlTreeNode[K, V], key K, value V) *avlTreeNode[K, V] {
	if n == nil {
		return &avlTreeNode[K, V]{key: key, value: value}
	}
	if key < n.key {
		n.left = insertAvlTreeNode(n.left, key, value)
	} else if key > n.key {
		n.right = insertAvlTreeNode(n.right, key, value)
	} else { // key == n.key
		n.value = value
	}
	return n.balance()
}

func removeAvlTreeNode[K orderedKey, V any](n *avlTreeNode[K, V], key K) *avlTreeNode[K, V] {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = removeAvlTreeNode(n.left, key)
	} else if key > n.key {
		n.right = removeAvlTreeNode(n.right, key)
	} else { // key = n.key
		left := n.left
		right := n.right
		n.left = nil
		n.right = nil
		if right == nil {
			return left
		}
		newRoot := right.findLeftmostChild()
		newRoot.right = right.removeLeftmostChild()
		newRoot.left = left

		return newRoot.balance()
	}
	return n.balance()
}

func listAvlTreeKeys[K orderedKey, V any](n *avlTreeNode[K, V], keys []K) []K {
	if n == nil {
		return keys
	}
	keys = listAvlTreeKeys(n.left, keys)
	keys = append(keys, n.key)
	keys = listAvlTreeKeys(n.right, keys)

	return keys
}

func listAvlTreeValues[K orderedKey, V any](n *avlTreeNode[K, V], values []V) []V {
	if n == nil {
		return values
	}
	values = listAvlTreeValues(n.left, values)
	values = append(values, n.value)
	values = listAvlTreeValues(n.right, values)

	return values
}
