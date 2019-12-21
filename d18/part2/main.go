package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"github.com/mm4tt/aoc2019/util"
	"log"
	"strings"
)

func main() {
	lines, err := util.ReadLines("d18/part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	b := newBoard()
	for line := range lines {
		b.add(line)
	}
	b.X, b.Y = len(b.rows[0]), len(b.rows)

	fmt.Println(b)
	deps := b.processDeps()
	fmt.Println("Deps before compaction")
	printDeps(deps)

	deps = compactDeps(b, deps)
	fmt.Println("Deps after compaction")
	printDeps(deps)

	dists := b.calculateDists()

	fmt.Println(findShortestPath(deps, dists))
}

const (
	EMPTY = '.'
	WALL  = '#'
)

func printDeps(deps map[rune][]rune) {
	for from, tos := range deps {
		fmt.Printf("%c -> %v\n", from, string(tos))
	}
}

type Board struct {
	X, Y  int
	start []util.Point2D
	rows  []string
	keys  map[rune]bool
	all   map[rune]util.Point2D
}

func newBoard() *Board {
	return &Board{
		keys: make(map[rune]bool),
		all:  make(map[rune]util.Point2D),
	}
}

func (b *Board) add(line string) {
	if line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}
	for x, c := range line {
		p := util.Point2D{X: x, Y: len(b.rows)}
		if '1' <= c && c <= '4' {
			b.start = append(b.start, p)
		}
		if isKey(c) {
			b.keys[c] = true
		}
		if isKeyOrDoor(c) || '1' <= c && c <= '4' {
			b.all[c] = p
		}
	}
	b.rows = append(b.rows, line)
	b.Y++
	if b.X == 0 {
		b.X = len(line)
	} else if b.X != len(line) {
		panic("Invalid board")
	}
}

func (b *Board) get(p util.Point2D) rune {
	return rune(b.rows[p.Y][p.X])
}

func (b *Board) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("X=%v, Y=%v, start=%v\n", b.X, b.Y, b.start))
	s.WriteString(fmt.Sprintf("# keys: %v\n", len(b.keys)))
	s.WriteString(fmt.Sprintf("# all: %v\n", len(b.all)))
	for _, l := range b.rows {
		s.WriteRune('\n')
		s.WriteString(l)
	}
	return s.String()
}

func (b *Board) neighbours(p util.Point2D) []util.Point2D {
	ret := []util.Point2D{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if util.Abs(x)+util.Abs(y) != 1 {
				continue
			}
			p1 := p.Add(util.Point2D{X: x, Y: y})
			if b.isOnBoard(p1) && b.get(p) != WALL {
				ret = append(ret, p1)
			}
		}
	}
	return ret
}

func (b *Board) isOnBoard(p util.Point2D) bool {
	return 0 <= p.X && p.X < b.X && 0 <= p.Y && p.Y < b.Y
}

func (b *Board) processDeps() map[rune][]rune {
	deps := make(map[rune][]rune)
	for k := range b.keys {
		d := toDoor(k)
		if _, ok := b.all[d]; ok {
			deps[k] = append(deps[k], d)
		}
	}

	for _, start := range b.start {
		s := stack.New()
		s.Push(start)
		type State int
		const (
			NOT_VISITED State = iota
			VISITING
			VISITED
		)
		state := map[util.Point2D]State{start: VISITING}
		parent := make(map[util.Point2D]rune)
		currentSymbol := '@'
		parent[start] = currentSymbol
	main:
		for ; s.Len() > 0; {
			p := s.Peek().(util.Point2D)
			if state[p] != VISITING {
				panic(fmt.Sprintf("Wrong state, expected VISITING, got: %v", state[p]))
			}
			for _, n := range b.neighbours(p) {
				t := b.get(n)
				switch state[n] {
				case VISITING:
					if parent[n] == parent[p] || parent[p] == b.get(n) {
						continue
					}
					panic("Something went wrong...")
				case NOT_VISITED:
					state[n] = VISITING
					parent[n] = currentSymbol
					if isKeyOrDoor(t) {
						if isKeyOrDoor(currentSymbol) {
							deps[currentSymbol] = append(deps[currentSymbol], t)
						}
						currentSymbol = t
					}
					s.Push(n)
					continue main
				}
			}
			state[p] = VISITED
			currentSymbol = parent[p]
			s.Pop()
		}
	}
	return deps

}

func (b *Board) calculateDists() map[rune]map[rune]int {
	dist := make(map[rune]map[rune]int)
	for k, p := range b.all {
		dist[k] = b.calculateDist(p)
	}
	return dist
}

func (b *Board) calculateDist(p util.Point2D) map[rune]int {
	dist, d := make(map[rune]int), make(map[util.Point2D]int)
	q := queue.New()
	q.Enqueue(p)
	d[p] = 0

	for ; q.Len() > 0; {
		p := q.Dequeue().(util.Point2D)
		if _, ok := b.all[b.get(p)]; ok {
			dist[b.get(p)] = d[p]
		}
		for _, n := range b.neighbours(p) {
			if _, ok := d[n]; ok {
				continue
			}
			d[n] = d[p] + 1
			q.Enqueue(n)
		}
	}
	return dist
}

func compactDeps(b *Board, deps map[rune][]rune) map[rune][]rune {
	for ; ; {
		hasAnythingChanged := false
		newDeps := make(map[rune][]rune)
		for from, tos := range deps {
			for _, to := range tos {
				if !isDoor(to) {
					newDeps[from] = append(newDeps[from], to)
					continue
				}
				hasAnythingChanged = true
				for _, to1 := range deps[to] {
					newDeps[from] = append(newDeps[from], to1)
				}
			}
		}
		deps = newDeps
		if !hasAnythingChanged {
			break
		}
	}

	newDeps := make(map[rune][]rune)
	for k := range b.keys {
		if len(deps[k]) > 0 {
			newDeps[k] = deps[k]
		}
	}
	return newDeps
}

func findShortestPath(deps map[rune][]rune, dists map[rune]map[rune]int) int {
	inDegree := make(map[rune]int)
	all, free := make(map[rune]bool), make(map[rune]bool)
	for from, tos := range deps {
		all[from] = true
		for _, to := range tos {
			all[to] = true
			inDegree[to]++
		}
	}
	for k := range all {
		if inDegree[k] == 0 {
			free[k] = true
		}
	}

	minDist := 1
	for _, tos := range dists {
		for _, d := range tos {
			minDist = util.Min(minDist, d)
		}
	}
	N, best := len(all), -1
	findShortestPathHelper(&best, 0, [4]rune{'1', '2', '3', '4'}, make([]rune, 0, N), free, inDegree, N, deps, dists, minDist)
	return best
}

func findShortestPathHelper(
	bestLength *int,
	currentLength int,
	current [4]rune,
	currentPath []rune,
	free map[rune]bool,
	inDegree map[rune]int,
	N int,
	deps map[rune][]rune,
	dists map[rune]map[rune]int,
	minDist int) {

	if len(free) == 0 {
		if *bestLength == -1 || currentLength < *bestLength {
			*bestLength = currentLength
			fmt.Printf("Best Path (%d): %s\n", currentLength, string(currentPath))
		}
		return
	}
	if *bestLength != -1 {
		minPossibleLength := currentLength + (N-len(currentPath)-1) * minDist
		if minPossibleLength >= *bestLength {
			return
		}
	}
	for _, c := range keyCopy(free) {
		quadrant := 0
		for ; quadrant < len(current); quadrant++ {
			if dists[current[quadrant]][c] != 0 {
				break
			}
		}
		if quadrant == len(current) {
			panic("Quadrant not found!")
		}

		delete(free, c)
		currentPath = append(currentPath, c)
		prevC := current[quadrant]
		current[quadrant] = c
		for _, to := range deps[c] {
			inDegree[to]--
			if inDegree[to] == 0 {
				free[to] = true
			}
		}

		findShortestPathHelper(bestLength, currentLength+dists[prevC][c], current, currentPath, free, inDegree, N, deps, dists, minDist)

		for _, to := range deps[c] {
			if inDegree[to] == 0 {
				delete(free, to)
			}
			inDegree[to]++
		}
		current[quadrant] = prevC
		currentPath = currentPath[:len(currentPath)-1]
		free[c] = true
	}
}

func toDoor(key rune) rune {
	return key + 'A' - 'a'
}

func isKey(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isDoor(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func isKeyOrDoor(r rune) bool {
	return isKey(r) || isDoor(r)
}

func printRuneSet(s map[rune]bool) {
	for k := range s {
		fmt.Print(string(k))
	}
	fmt.Println()
}

func keyCopy(m map[rune]bool) []rune {
	ret := make([]rune, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	//sort.Slice(ret, func(i, j int) bool {
	//	return ret[i] < ret[j]
	//})
	return ret
}
