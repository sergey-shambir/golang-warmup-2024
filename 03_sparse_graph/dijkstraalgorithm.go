package main

import (
	"container/heap"
	"math"
)

type dijkstraAlgorithmState struct {
	distances []float64
	minHeap   *minDistanceHeap
}

func newDijkstraAlgorithmState(vertexCount uint) *dijkstraAlgorithmState {
	state := &dijkstraAlgorithmState{
		distances: make([]float64, vertexCount),
		minHeap:   &minDistanceHeap{},
	}

	infinity := math.Inf(1)
	for i := range vertexCount {
		state.distances[i] = infinity
	}

	return state
}

func (s *dijkstraAlgorithmState) distance(vertex uint) float64 {
	return s.distances[vertex]
}

func (s *dijkstraAlgorithmState) updateDistance(vertex uint, distance float64) {
	s.distances[vertex] = distance
	heap.Push(s.minHeap, distanceTo{distance: distance, vertex: vertex})
}

func (s *dijkstraAlgorithmState) isFinished() bool {
	return len(*s.minHeap) == 0
}

func (s *dijkstraAlgorithmState) popMinDistance() distanceTo {
	return heap.Pop(s.minHeap).(distanceTo)
}
