package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"github.com/mm4tt/aoc2019/util"
	"log"
	"math"
	"sort"
	"strings"
)

func main() {
	lines, err := util.ReadLines("d18/input.txt")
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
	deps = reverseDeps(deps)

	dists := b.calculateDists()
	printDists(dists)
	distSums := distSums(dists)
	fmt.Println("Dist sums:", distSums)

	fmt.Println("\nSolution:")
	fmt.Println(findShortestPath(b, deps, dists, distSums))
}

type Node struct {
	sortedPath string
	current    rune
}

func (n *Node) Current() rune {
	return n.current
}

const INF = math.MaxInt64

func findShortestPath(b *Board, deps map[rune][]rune, dists map[rune]map[rune]int, distSums []int) (int, string) {
	dist := make(map[Node]int)
	h := func(n Node) int {
		return distSums[(len(b.keys) + 1 - len(n.sortedPath))]
	}

	priority := func(x interface{}) int {
		if n, ok := x.(Node); ok {
			if d, ok := dist[n]; ok {
				return d + h(n)
			}
		}
		panic(fmt.Sprintf("Wrong object: %v", x))
	}
	isTerminal := func(n Node) bool {
		return len(n.sortedPath) == len(b.keys)+1 // keys + '@'
	}
	canMoveTo := func(keys map[rune]bool, to rune) bool {
		for _, k := range deps[to] {
			if !keys[k] {
				return false
			}
		}
		return true
	}

	pq := util.NewPriorityQueue(priority)
	start := Node{sortedPath: "@", current: '@'}
	dist[start] = 0
	pq.Push(start)

	for ; ; {
		n := pq.Pop().(Node)
		//fmt.Println("Processing", n, dist[n] + h(n))
		if isTerminal(n) {
			return dist[n], n.sortedPath
		}
		c := n.Current()
		keys := make(map[rune]bool)
		for _, r := range n.sortedPath {
			keys[r] = true
		}

		for t, e := range dists[c] {
			if keys[t] || !canMoveTo(keys, t) {
				continue
			}
			newPath := make([]rune, len(n.sortedPath) +1)
			for i, c := range n.sortedPath {
				newPath[i+1] = c
			}
			newPath[0] = t
			for i := 0; i + 1 < len(newPath); i++ {
				if newPath[i] < newPath[i+1] {
					break
				}
				newPath[i], newPath[i+1] = newPath[i+1], newPath[i]
			}

			n2 := Node{
				current: t,
				sortedPath: string(newPath),
			}
			d, ok := dist[n2]
			if !ok {
				d = INF
			}
			if d > dist[n]+e {
				dist[n2] = dist[n] + e
			}
			if ok {
				pq.Update(n2)
			} else {
				pq.Push(n2)
			}
		}
	}
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
	start util.Point2D
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
		if c == '@' {
			b.start = p
		}
		if isKey(c) {
			b.keys[c] = true
		}
		if isKeyOrDoor(c) || c == '@' {
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

	s := stack.New()
	s.Push(b.start)
	type State int
	const (
		NOT_VISITED State = iota
		VISITING
		VISITED
	)
	state := map[util.Point2D]State{b.start: VISITING}
	parent := make(map[util.Point2D]rune)
	currentSymbol := '@'
	parent[b.start] = currentSymbol
main:
	for ; s.Len() > 0; {
		p := s.Peek().(util.Point2D)
		//fmt.Println("Visiting ", p)
		if state[p] != VISITING {
			panic(fmt.Sprintf("Wrong state, expected VISITING, got: %v", state[p]))
		}
		for _, n := range b.neighbours(p) {
			t := b.get(n)
			//fmt.Println("\tNeighbour is: ", n)
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
	return deps
}

func (b *Board) calculateDists() map[rune]map[rune]int {
	dist := make(map[rune]map[rune]int)
	for k, p := range b.all {
		if isDoor(k) {
			continue
		}
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
		if _, ok := b.keys[b.get(p)]; ok {
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

func reverseDeps(deps map[rune][]rune) map[rune][]rune {
	ret := make(map[rune][]rune)
	for from, tos := range deps {
		for _, to := range tos {
			ret[to] = append(ret[to], from)
		}
	}
	return ret
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

func printDists(dists map[rune]map[rune]int) {
	for from, dist := range dists {
		fmt.Println(string(from))
		for a, b := range dist {
			fmt.Printf("\t->%c: %d\n", a, b)
		}
	}
}

func distSums(dists map[rune]map[rune]int) []int {
	ds := []int{}
	for a, dist := range dists {
		for b, d := range dist {
			if a == b {
				continue
			}
			ds = append(ds, d)
		}
	}
	sort.Ints(ds)
	for i := range ds {
		if i == 0 {
			continue
		}
		ds[i] += ds[i-1]
	}
	return ds
}
