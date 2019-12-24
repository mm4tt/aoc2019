package main

import (
	"fmt"
	"github.com/mm4tt/aoc2019/util"
	"log"
)

func main() {
	lines, err := util.ReadLines("d24/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	b := newBoard()
	y := 0
	for line := range lines {
		for x, c := range line {
			b.Set(x, y, c)
		}
		y++
	}

	b.Print()

	seen := make(map[Board]bool)
	for ; !seen[b]; {
		seen[b] = true
		b = b.Next()
	}
	fmt.Printf("\nAfter %d iterations\n", len(seen))
	b.Print()
	fmt.Println("BiodiversityRating:", b.BiodiversityRating())
}

const (
	N     = 5
	BUG   = '#'
	EMPTY = '.'
)

type Board struct {
	s string
}

func newBoard() Board {
	return Board{string(make([]rune, N*N))}
}

func (b *Board) Next() Board {
	b1 := newBoard()
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			b1.Set(x, y, b.Get(x, y))
			nAdjacentBugs := b.numSurroundingBugs(x, y)
			switch b.Get(x, y) {
			case BUG:
				if nAdjacentBugs != 1 {
					b1.Set(x, y, EMPTY)
				}
			case EMPTY:
				if nAdjacentBugs == 1 || nAdjacentBugs == 2 {
					b1.Set(x, y, BUG)
				}
			}
		}
	}
	return b1
}

func (b *Board) BiodiversityRating() int {
	r := 0
	c := 1
	for i := 0; i < N*N; i++ {
		if b.s[i] == BUG {
			r += c
		}
		c *= 2
	}
	return r
}

func (b *Board) Get(x, y int) rune {
	return rune(b.s[b.index(x, y)])
}

func (b *Board) Set(x, y int, v rune) {
	i := b.index(x, y)
	r := []rune(b.s)
	r[i] = v
	b.s = string(r)
}

func (b *Board) index(x, y int) int {
	if !b.isOnBoard(x, y) {
		panic(fmt.Sprintf("index out of bounds: x=%d, y=%d", x, y))
	}
	return N*y + x
}

func (b *Board) Print() {
	for i := 0; i < N; i++ {
		fmt.Println(b.s[i : i+N])
	}
}

func (b *Board) numSurroundingBugs(pX, pY int) int {
	ret := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if util.Abs(x)+util.Abs(y) != 1 {
				continue
			}
			if b.isOnBoard(pX+x, pY+y) && b.Get(pX+x, pY+y) == BUG {
				ret++
			}
		}
	}
	return ret
}

func (b *Board) isOnBoard(x, y int) bool {
	return 0 <= x && x < N && 0 <= y && y < N
}
