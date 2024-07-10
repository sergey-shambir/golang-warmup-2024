package main

type Point struct {
	x, y int
}

type finder struct {
	cells         [][]int
	height, width int
	visited       [][]bool
}

func newFinder(cells [][]int) *finder {
	width := len(cells[0])
	height := len(cells)

	visited := make([][]bool, height)
	for i := range height {
		visited[i] = make([]bool, width)
	}

	return &finder{
		cells:   cells,
		height:  len(cells),
		width:   len(cells[0]),
		visited: visited,
	}
}

func (f *finder) canAchievePoint(from, to Point) bool {
	if !f.isValidPoint(from) || !f.isValidPoint(to) {
		return false
	}
	if from == to {
		return true
	}
	if f.visited[from.y][from.x] {
		return false
	}
	f.visited[from.y][from.x] = true

	jumpLength := f.cells[from.y][from.x]
	left := Point{from.x - jumpLength, from.y}
	right := Point{from.x + jumpLength, from.y}
	top := Point{from.x, from.y - jumpLength}
	bottom := Point{from.x, from.y + jumpLength}

	return f.canAchievePoint(left, to) ||
		f.canAchievePoint(right, to) ||
		f.canAchievePoint(top, to) ||
		f.canAchievePoint(bottom, to)
}

func (f *finder) isValidPoint(p Point) bool {
	return p.x >= 0 && p.x < f.width && p.y >= 0 && p.y < f.height
}

func CanAchievePoint(cells [][]int, from, to Point) bool {
	f := newFinder(cells)
	return f.canAchievePoint(from, to)
}
