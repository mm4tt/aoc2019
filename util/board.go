package util

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"strings"
)

type Board struct {
	X, Y int
	rows []string
}

func NewBoard(lines []string) *Board {
	b := &Board{}
	b.rows = lines
	b.preprocess()
	return b
}

func (b *Board) preprocess() {
	b.Y = len(b.rows)
	// Compute b.X and extend rows if needed
	b.X = 0
	for _, line := range b.rows {
		b.X = Max(b.X, len(line))
	}
	for i, line := range b.rows {
		if len(line) == b.X {
			continue
		}
		b.rows[i] += string(make([]rune, b.X-len(line)))
	}
}

func (b *Board) Get(p Point2D) rune {
	return rune(b.rows[p.Y][p.X])
}

func (b *Board) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("X=%v, Y=%v\n", b.X, b.Y))
	for _, l := range b.rows {
		s.WriteRune('\n')
		s.WriteString(l)
	}
	return s.String()
}

func (b *Board) Neighbours(p Point2D, matchers ...func(rune) bool) []Point2D {
	ret := []Point2D{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if Abs(x)+Abs(y) != 1 {
				continue
			}
			p1 := p.Add(Point2D{X: x, Y: y})
			if b.IsOnBoard(p1) && matches(b.Get(p1), matchers) {
				ret = append(ret, p1)
			}
		}
	}
	return ret
}

func (b *Board) IsOnBoard(p Point2D) bool {
	return 0 <= p.X && p.X < b.X && 0 <= p.Y && p.Y < b.Y
}

func (b *Board) CalculateDist(p Point2D, allowed ...func(rune) bool) map[Point2D]int {
	dist := make(map[Point2D]int)
	q := queue.New()
	q.Enqueue(p)
	dist[p] = 0
	for ; q.Len() > 0; {
		p := q.Dequeue().(Point2D)
		for _, n := range b.Neighbours(p, allowed...) {
			if _, ok := dist[n]; ok {
				continue
			}
			dist[n] = dist[p] + 1
			q.Enqueue(n)
		}
	}
	return dist
}

func matches(r rune, matchers []func(rune) bool) bool {
	if len(matchers) == 0 {
		return true
	}
	for _, m := range matchers {
		if m(r) {
			return true
		}
	}
	return false
}
