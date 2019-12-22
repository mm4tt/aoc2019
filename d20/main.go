package main

import (
	"fmt"
	"github.com/mm4tt/aoc2019/util"
	"log"
)

func main() {
	lines, err := util.ReadAllLines("d20/input0.txt")
	if err != nil {
		log.Fatal(err)
	}
	b := newBoard(lines)
	fmt.Println(b)
	fmt.Println(b.tagToPoint)
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
	}
	b.preprocess()
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
			if _, ok := b.tagToPoint[tag]; ok {
				tag = tag + "1"
			}
			if _, ok := b.tagToPoint[tag]; ok {
				panic("Tag already exist!")
			}
			// Find '.' char next to either first or second letter.
			n = append(b.Neighbours(p, is(EMPTY)), b.Neighbours(a2P, is(EMPTY))...)
			if len(n) != 1 {
				panic("Couldn't find . next to portal")
			}
			p1 := n[0]
			// Store tag
			b.tagToPoint[tag] = p1
			b.pointToTag[p1] = tag
		}
	}
}

func isLetter(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func is(r rune) func(rune) bool {
	return func(a rune) bool {
		return r == a
	}
}
