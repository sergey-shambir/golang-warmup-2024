package main

/**
 * Разреженный граф с последовательными индексами вершин
 */
type SparseGraph struct {
	verticies []FlatMap[uint, float64]
}

func NewSparseGraph(vertexCount uint) *SparseGraph {
	return &SparseGraph{
		verticies: make([]FlatMap[uint, float64], vertexCount),
	}
}

func (g *SparseGraph) VertexCount() uint {
	return uint(len(g.verticies))
}

func (g *SparseGraph) HasEdge(fromVertex, toVertex uint) bool {
	return g.verticies[fromVertex].Contains(toVertex)
}

func (g *SparseGraph) EdgeWeight(fromVertex, toVertex uint) float64 {
	return g.verticies[fromVertex].Get(toVertex)
}

func (g *SparseGraph) AdjacencyList(fromVertex uint) FlatMap[uint, float64] {
	return g.verticies[fromVertex]
}

func (g *SparseGraph) AddDirectedAdge(fromVertex, toVertex uint, weight float64) {
	g.verticies[fromVertex].Add(toVertex, weight)
}

func (g *SparseGraph) AddUndirectedEdge(fromVertex, toVertex uint, weight float64) {
	g.verticies[fromVertex].Add(toVertex, weight)
	g.verticies[toVertex].Add(fromVertex, weight)
}

func (g *SparseGraph) FindDistancesFrom(fromVertex uint) []float64 {
	state := newDijkstraAlgorithmState(g.VertexCount())
	state.updateDistance(fromVertex, 0.0)

	for !state.isFinished() {
		distanceTo := state.popMinDistance()
		currentDistance := state.distance(distanceTo.vertex)
		if distanceTo.distance > currentDistance {
			continue // Ignore outdated distances from min-heap (priority queue)
		}
		for _, edge := range g.AdjacencyList(distanceTo.vertex).Items() {
			newDistance := currentDistance + edge.value
			if newDistance < state.distance(edge.key) {
				state.updateDistance(edge.key, newDistance)
			}
		}
	}

	return state.distances
}
