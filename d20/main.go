package main

import (
	"fmt"
	"github.com/mm4tt/aoc2019/util"
	"log"
	"math"
)

const INF = math.MaxInt64

func main() {
	lines, err := util.ReadAllLines("d20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b := newBoard(lines)
	fmt.Println(b)
	fmt.Println(b.tagToPoint)
	fmt.Println(b.dist)

	dist := dijkstra(b.dist, "AA")
	fmt.Println(dist)
	fmt.Println(dist["ZZ"])

	// Part2
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println(dijkstra2(b.dist))
}

func dijkstra(edges map[string]map[string]int, start string) map[string]int {
	dist := make(map[string]int)
	toProcess := make(map[string]bool)
	for v := range edges {
		toProcess[v] = true
		dist[v] = INF
	}
	dist[start] = 0

	getClosest := func() string {
		min, best := INF, ""
		for v := range toProcess {
			if dist[v] < min {
				min, best = dist[v], v
			}
		}
		if min == INF {
			panic("Something went wrong")
		}
		return best
	}

	for ; len(toProcess) > 0; {
		v := getClosest()
		delete(toProcess, v)
		for v2, d := range edges[v] {
			if dist[v2] > dist[v]+d {
				dist[v2] = dist[v] + d
			}
		}
	}
	return dist
}

func dijkstra2(edges map[string]map[string]int) int {
	type Node struct {
		tag   string
		level int
	}

	start, end := Node{"AA", 0}, Node{"ZZ", 0}

	d := make(map[Node]int)
	toProcess := make(map[Node]bool)
	dist := func(v Node) int {
		if d, ok := d[v]; ok {
			return d
		}
		return INF
	}
	d[start] = 0
	toProcess[start] = true

	getClosest := func() Node {
		min, best := INF, Node{}
		for v := range toProcess {
			if d := dist(v); d < min {
				min, best = d, v
			}
		}
		if min == INF {
			panic("Something went wrong")
		}
		return best
	}

	for ; ; {
		v := getClosest()
		if v == end {
			return dist(v)
		}
		delete(toProcess, v)
		for t, e := range edges[v.tag] {
			l := v.level
			if arePair(v.tag, t) {
				if l == 0 && isOuter(v.tag) {
					continue
				}
				if isInner(v.tag) {
					l++
				} else {
					l--
				}
				if l < 0 {
					continue
				}
			}
			v2 := Node{tag: t, level: l}
			if dist(v2) > dist(v)+e {
				d[v2] = dist(v) + e
				toProcess[v2] = true
			}
		}
	}
}

func arePair(t1, t2 string) bool {
	return t1!=t2 && t1[:2] == t2[:2]
}

func isInner(t string) bool {
	return t[3] == 'I'
}

func isOuter(t string) bool {
	return t[3] == 'O'
}

const (
	EMPTY = '.'
)

type Board struct {
	*util.Board
	pointToTag map[util.Point2D]string
	tagToPoint map[string]util.Point2D
	dist       map[string]map[string]int
}

func newBoard(lines []string) *Board {
	b := &Board{
		Board:      util.NewBoard(lines),
		pointToTag: make(map[util.Point2D]string),
		tagToPoint: make(map[string]util.Point2D),
		dist:       make(map[string]map[string]int),
	}
	b.preprocess()
	b.initializeDist()
	return b
}

func (b *Board) preprocess() {
	processedLetters := make(map[util.Point2D]bool)
	isProcessed := func(p util.Point2D) bool { return processedLetters[p] }
	markProcessed := func(p util.Point2D) { processedLetters[p] = true }

	p := util.Point2D{}
	for p.Y = 0; p.Y < b.Y; p.Y++ {
		for p.X = 0; p.X < b.X; p.X++ {
			a1 := b.Get(p)
			if !isLetter(a1) || isProcessed(p) {
				continue
			}
			// Find a2, second letter
			n := b.Neighbours(p, isLetter)
			if len(n) != 1 {
				panic("Couldn't find second letter")
			}
			a2P := n[0]
			a2 := b.Get(a2P)
			markProcessed(a2P)
			// Compute tag
			tag := string([]rune{a1, a2})
			// Find '.' char next to either first or second letter.
			n = append(b.Neighbours(p, is(EMPTY)), b.Neighbours(a2P, is(EMPTY))...)
			if len(n) != 1 {
				panic("Couldn't find . next to portal")
			}
			p1 := n[0]

			// Is it inner or outer?
			if tag != "AA" && tag != "ZZ" {
				if p1.X == 2 || p1.X == b.X-3 || p1.Y == 2 || p1.Y == b.Y-3 {
					tag += "_O" // outer
				} else {
					tag += "_I" // inner
				}
			}
			if _, ok := b.tagToPoint[tag]; ok {
				panic("Tag already exist!")
			}
			// Store tag
			b.tagToPoint[tag] = p1
			b.pointToTag[p1] = tag
		}
	}
}

func (b *Board) initializeDist() {
	for tag := range b.tagToPoint {
		b.dist[tag] = b.calculateDist(tag)
	}
	for tag := range b.tagToPoint {
		if len(tag) != 4 {
			continue
		}
		tag1, tag2 := tag[:2]+"_I", tag[:2]+"_O"
		b.dist[tag2][tag1] = 1
		b.dist[tag1][tag2] = 1
	}
}

func (b *Board) calculateDist(tag string) map[string]int {
	ret := make(map[string]int)
	dist := b.CalculateDist(b.tagToPoint[tag], is(EMPTY))
	for tag, p := range b.tagToPoint {
		if d, ok := dist[p]; ok {
			ret[tag] = d
		}
	}
	return ret
}

func isLetter(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func is(r rune) func(rune) bool {
	return func(a rune) bool {
		return r == a
	}
}
