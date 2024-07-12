package main

type point struct {
	x, y int32
}

func (p point) toInt64() int64 {
	return int64(p.x)<<32 + int64(p.y)
}

func (p point) digitsSum() int {
	return digitsSum(p.x) + digitsSum(p.y)
}

func digitsSum(num int32) int {
	if num < 0 {
		num = -num
	}
	var sum int
	for num > 0 {
		sum += int(num % 10)
		num = num / 10
	}
	return sum
}

func countAchievableCells(p point, maxDigitsSum int) int {
	visited := make(map[int64]bool)
	toBeVisited := []point{p}

	for len(toBeVisited) > 0 {
		p := toBeVisited[len(toBeVisited)-1]
		toBeVisited = toBeVisited[0 : len(toBeVisited)-1]

		if p.digitsSum() > maxDigitsSum || visited[p.toInt64()] {
			continue
		}
		visited[p.toInt64()] = true

		toBeVisited = append(toBeVisited, point{p.x - 1, p.y}, point{p.x + 1, p.y}, point{p.x, p.y - 1}, point{p.x, p.y + 1})
	}

	return len(visited)
}
