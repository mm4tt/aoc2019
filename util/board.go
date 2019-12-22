package util

import (
	"fmt"
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
