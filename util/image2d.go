package util

import "fmt"

func Draw(m map[Point2D]int) {
	xMin, xMax, yMin, yMax := getBoundaries(m)
	n := xMax - xMin + 1
	line := make([]rune, n)
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			p := Point2D{x, y}
			if  m[p] == 0 {
				line[x - xMin] = ' '
			} else {
				line[x - xMin] = '#'
			}
		}
		fmt.Println(string(line))
	}
}

func getBoundaries(m map[Point2D]int) (xMin, xMax, yMin, yMax int) {
	firstIteration := true
	for p := range m {
		if firstIteration {
			xMin, xMax, yMin, yMax = p.X, p.X, p.Y, p.Y
			firstIteration = false
			continue
		}
		if xMin > p.X {
			xMin = p.X
		}
		if xMax < p.X {
			xMax = p.X
		}
		if yMin > p.Y {
			yMin = p.Y
		}
		if yMax < p.Y {
			yMax = p.Y
		}
	}
	return xMin - 1, xMax + 1, yMin - 1, yMax + 1
}
