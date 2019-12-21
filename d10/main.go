package main

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/mm4tt/aoc2019/util"
)

func main() {
	inputLines, err := util.ReadLines("d10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	asteroids := []asteroid{}
	y := 0
	for line := range inputLines {
		for x, c := range line {
			if c == '#' {
				asteroids = append(asteroids, asteroid{Point2D: util.Point2D{x, y}})
			}
		}
		y++
	}

	// Part 1
	max, best := 0, asteroid{}
	for i, a := range asteroids {
		set := make(map[float64]bool)
		for j, b := range asteroids {
			if i == j {
				continue
			}
			set[computeRelativePolarCoord(a, b).a] = true
		}
		if max < len(set) {
			max = len(set)
			best = asteroids[i]
		}
	}
	fmt.Println(max, best)

	// Part 2
	angleToAsteroid := make(map[float64]asteroid)
	for _, a := range asteroids {
		if a == best {
			continue
		}
		a.polarCoord = computeRelativePolarCoord(best, a)
		c, ok := angleToAsteroid[a.a]
		if !ok || a.r < c.r {
			angleToAsteroid[a.a] = a
		}
	}
	angles := []float64{}
	for a := range angleToAsteroid {
		angles = append(angles, a)
	}
	sort.Float64s(angles)
	fmt.Println(len(angles))

	a := angleToAsteroid[angles[199]]
	fmt.Println(a.X*100 + a.Y)
}

type asteroid struct {
	util.Point2D
	polarCoord
}

type polarCoord struct {
	a float64
	r float64
}

func computeRelativePolarCoord(base, p asteroid) polarCoord {
	return computePolarCoord(p.X-base.X, base.Y-p.Y)
}

func computePolarCoord(xI, yI int) polarCoord {
	x, y := float64(xI), float64(yI)
	r := math.Sqrt(x*x + y*y)
	a := -math.Atan2(y, x) + math.Pi/2
	if a < 0 {
		a += 2 * math.Pi
	}
	return polarCoord{a, r}
}

func testPolarCoord() {
	fmt.Println(
		computePolarCoord(0, 2),
		computePolarCoord(2, 0),
		computePolarCoord(0, -2),
		computePolarCoord(-2, 0),
		computePolarCoord(-3, 4),
	)
}
