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
	init := b

	b.Print()

	seen := make(map[Board]bool)
	for ; !seen[b]; {
		seen[b] = true
		b = b.Next()
	}
	fmt.Printf("\nAfter %d iterations\n", len(seen))
	b.Print()
	fmt.Println("BiodiversityRating:", b.BiodiversityRating())

	fmt.Println("\nPart2")
	init.Print()
	r := NewRecursiveBoards(init)
	fmt.Println("Init # bugs: ", r.CountBugs())

	for i := 0; i < 200; i++ {
		r.Iterate()
	}
	fmt.Println("After 200 minutes # bugs: ", r.CountBugs())
}

const (
	N     = 5
	BUG   = '#'
	EMPTY = '.'
)

type RecursiveBoards struct {
	boards map[int]Board

	iterations int
}

func NewRecursiveBoards(init Board) *RecursiveBoards {
	return &RecursiveBoards{
		boards:     map[int]Board{0: init},
		iterations: 0,
	}
}

func (r *RecursiveBoards) CountBugs() int {
	count := 0
	for _, b := range r.boards {
		count += b.NumBugs()
	}
	return count
}

func (r *RecursiveBoards) Iterate() {
	r.iterations++
	r.boards[r.iterations] = newBoard()
	r.boards[-r.iterations] = newBoard()

	newBoards := make(map[int]Board)
	for level, b := range r.boards {
		newB := newBoard()

		for y := 0; y < N; y++ {
			for x := 0; x < N; x++ {
				if x == 2 && y == 2 {
					continue
				}
				newB.Set(x, y, b.Get(x, y))
				nAdjacentBugs := r.numSurroundingBugs(level, x, y)
				switch b.Get(x, y) {
				case BUG:
					if nAdjacentBugs != 1 {
						newB.Set(x, y, EMPTY)
					}
				case EMPTY:
					if nAdjacentBugs == 1 || nAdjacentBugs == 2 {
						newB.Set(x, y, BUG)
					}
				}
			}
		}

		newBoards[level] = newB
	}
	r.boards = newBoards
}

func (r *RecursiveBoards) numSurroundingBugs(level, x, y int) int {
	ret := 0
	b := r.boards[level]
	for xS := -1; xS <= 1; xS++ {
		for yS := -1; yS <= 1; yS++ {
			if util.Abs(xS)+util.Abs(yS) != 1 {
				continue
			}
			x1, y1 := x+xS, y+yS
			if x1 == 2 && y1 == 2 {
				continue
			}
			if b.isOnBoard(x1, y1) {
				ret += b.nBugs(x1, y1)
			}
		}
	}
	if level != -r.iterations {
		if x == 0 {
			ret += r.boards[level-1].nBugs(1, 2)
		}
		if y == 0 {
			ret += r.boards[level-1].nBugs(2, 1)
		}
		if x == N-1 {
			ret += r.boards[level-1].nBugs(3, 2)
		}
		if y == N-1 {
			ret += r.boards[level-1].nBugs(2, 3)
		}
	}

	if level != r.iterations {
		if x == 1 && y == 2 {
			ret += r.boards[level+1].nBugsInCol(0)
		}
		if x == 2 && y == 1 {
			ret += r.boards[level+1].nBugsInRow(0)
		}
		if x == 3 && y == 2 {
			ret += r.boards[level+1].nBugsInCol(N - 1)
		}
		if x == 2 && y == 3 {
			ret += r.boards[level+1].nBugsInRow(N - 1)
		}
	}

	return ret
}

type Board struct {
	s string
}

func newBoard() Board {
	b := Board{string(make([]rune, N*N))}
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			b.Set(x, y, EMPTY)
		}
	}
	return b
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
		fmt.Println(b.s[i*N : (i+1)*N])
	}
}

func (b *Board) NumBugs() int {
	count := 0
	for i := 0; i < N*N; i++ {
		if b.s[i] == BUG {
			count++
		}
	}
	return count
}

func (b *Board) numSurroundingBugs(pX, pY int) int {
	ret := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if util.Abs(x)+util.Abs(y) != 1 {
				continue
			}
			if b.isOnBoard(pX+x, pY+y) {
				ret += b.nBugs(pX+x, pY+y)
			}
		}
	}
	return ret
}

func (b Board) nBugs(x, y int) int {
	if b.Get(x, y) == BUG {
		return 1
	}
	return 0
}

func (b Board) nBugsInCol(x int) int {
	count := 0
	for y := 0; y < N; y++ {
		count += b.nBugs(x, y)
	}
	return count
}

func (b Board) nBugsInRow(y int) int {
	count := 0
	for x := 0; x < N; x++ {
		count += b.nBugs(x, y)
	}
	return count
}

func (b *Board) isOnBoard(x, y int) bool {
	return 0 <= x && x < N && 0 <= y && y < N
}
