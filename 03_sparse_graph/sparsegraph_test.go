package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparseGraph(t *testing.T) {
	g := NewSparseGraph(4)
	g.AddUndirectedEdge(0, 3, 240.0)
	g.AddUndirectedEdge(0, 2, 200.0)
	g.AddUndirectedEdge(2, 3, 100.0)
	g.AddUndirectedEdge(1, 3, 300.0)

	assert.True(t, !g.HasEdge(0, 1))
	assert.True(t, !g.HasEdge(1, 1))
	assert.Equal(t, g.EdgeWeight(0, 0), 0.0)
	assert.Equal(t, g.EdgeWeight(0, 1), 0.0)

	assert.True(t, g.HasEdge(0, 3))
	assert.True(t, g.HasEdge(3, 0))
	assert.True(t, g.HasEdge(0, 2))
	assert.True(t, g.HasEdge(2, 0))
	assert.True(t, g.HasEdge(1, 3))
	assert.True(t, g.HasEdge(3, 1))

	assert.Equal(t, g.EdgeWeight(0, 3), 240.0)
	assert.Equal(t, g.EdgeWeight(3, 0), 240.0)
	assert.Equal(t, g.EdgeWeight(0, 2), 200.0)
	assert.Equal(t, g.EdgeWeight(2, 0), 200.0)
	assert.Equal(t, g.EdgeWeight(2, 3), 100.0)
	assert.Equal(t, g.EdgeWeight(3, 2), 100.0)
	assert.Equal(t, g.EdgeWeight(1, 3), 300.0)
	assert.Equal(t, g.EdgeWeight(3, 1), 300.0)
}

func TestFindDistanceForSmallGraph(t *testing.T) {
	g := NewSparseGraph(4)
	g.AddUndirectedEdge(0, 3, 240.0)
	g.AddUndirectedEdge(0, 2, 200.0)
	g.AddUndirectedEdge(2, 3, 100.0)
	g.AddUndirectedEdge(1, 3, 300.0)

	assert.Equal(t, g.FindDistancesFrom(0), []float64{0.0, 540.0, 200.0, 240.0})
	assert.Equal(t, g.FindDistancesFrom(1), []float64{540.0, 0.0, 400.0, 300.0})
	assert.Equal(t, g.FindDistancesFrom(2), []float64{200.0, 400.0, 0.0, 100.0})
	assert.Equal(t, g.FindDistancesFrom(3), []float64{240.0, 300.0, 100.0, 0.0})
}

func TestFindDistanceForAnotherSmallGraph(t *testing.T) {
	g := NewSparseGraph(4)
	g.AddUndirectedEdge(0, 3, 240.0)
	g.AddUndirectedEdge(0, 2, 200.0)
	g.AddUndirectedEdge(2, 3, 20.0)
	g.AddUndirectedEdge(1, 3, 300.0)

	assert.Equal(t, g.FindDistancesFrom(0), []float64{0.0, 520.0, 200.0, 220.0})
	assert.Equal(t, g.FindDistancesFrom(1), []float64{520.0, 0.0, 320.0, 300.0})
	assert.Equal(t, g.FindDistancesFrom(2), []float64{200.0, 320.0, 0.0, 20.0})
	assert.Equal(t, g.FindDistancesFrom(3), []float64{220.0, 300.0, 20.0, 0.0})
}

func TestFindDistanceForGraphWithOneVertex(t *testing.T) {
	g := NewSparseGraph(4)
	g.AddUndirectedEdge(0, 3, 240.0)
	g.AddUndirectedEdge(0, 2, 200.0)
	g.AddUndirectedEdge(2, 3, 20.0)
	g.AddUndirectedEdge(1, 3, 300.0)

	assert.Equal(t, g.FindDistancesFrom(0), []float64{0.0, 520.0, 200.0, 220.0})
	assert.Equal(t, g.FindDistancesFrom(1), []float64{520.0, 0.0, 320.0, 300.0})
	assert.Equal(t, g.FindDistancesFrom(2), []float64{200.0, 320.0, 0.0, 20.0})
	assert.Equal(t, g.FindDistancesFrom(3), []float64{220.0, 300.0, 20.0, 0.0})
}
