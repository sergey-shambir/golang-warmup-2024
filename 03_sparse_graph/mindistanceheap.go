package main

type distanceTo struct {
	distance float64
	vertex   uint
}

type minDistanceHeap []distanceTo

func (h minDistanceHeap) Len() int {
	return len(h)
}

func (h minDistanceHeap) Less(i, j int) bool {
	return h[i].distance < h[i].distance
}

func (h minDistanceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minDistanceHeap) Push(value any) {
	*h = append(*h, value.(distanceTo))
}

func (h *minDistanceHeap) Pop() any {
	oldHeap := *h
	size := len(oldHeap)
	result := oldHeap[size-1]
	*h = oldHeap[0 : size-1]

	return result
}
