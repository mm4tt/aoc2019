package main

import (
	"fmt"
	"github.com/mm4tt/aoc2019/util"
	"strconv"
	"strings"
)

func main() {
	N := 10007
	//N = 10

	g := NewGen2(N)
	transformations := loadTransformations(N)
	for _, tr := range transformations {
		g = tr.Apply(g)
	}

	g.Print(10)
	for i := 0; i < N; i++ {
		if g.Get(i) == 2019 {
			fmt.Println(i)
			break
		}
	}

	// Part2
	N = 119315717514047
	transformations = loadTransformations(N)
	K := len(transformations)

	M := 101741582076661 * K

	//M = 1000000; bruteCheck(N, M, transformations)

	kTr := IdentityTransformation(N)
	for _, t := range transformations {
		kTr = kTr.Compose(t)
	}

	tr := IdentityTransformation(N)
	for ; M >= K; {
		k := K
		cTr := kTr
		for ; 2*k <= M; k *= 2 {
			cTr = cTr.Compose(cTr)
		}
		M -= k
		fmt.Println("\t", "M", M, "k", k)
		tr = tr.Compose(cTr)
	}

	for i, t := range transformations {
		if i == M {
			break
		}
		tr = tr.Compose(t)
	}

	g = NewGen2(N)
	g = tr.Apply(g)
	g.Print(10)
	fmt.Println(g.Get(2020))
}

// 84599491118376 -- too high!
// 45994781087946 -- too high!
// 7757787935983

func bruteCheck(N, M int, transformations []Transformation) {
	g := NewGen2(N)
	for i := 0; i < M; {
		for _, tr := range transformations {
			g = tr.Apply(g)
			i++
			if i == M {
				break
			}
		}
	}
	fmt.Println("M: ", M)
	fmt.Println("BRUTE FORCE:")
	g.Print(10)
	fmt.Println(g.Get(2020))
	fmt.Println()
}

func loadTransformations(N int) []Transformation {
	lines, err := util.ReadLines("d22/input.txt")
	if err != nil {
		panic(err)
	}
	ret := []Transformation{}
	for line := range lines {
		if strings.HasPrefix(line, "cut") {
			n := atoi(line[4:])
			ret = append(ret, cutNCards(N, n))
		} else if strings.HasPrefix(line, "deal into new stack") {
			ret = append(ret, dealIntoNewStack(N))
		} else if strings.HasPrefix(line, "deal with increment ") {
			n := atoi(line[len("deal with increment "):])
			ret = append(ret, dealWithIncrementN(N, n))
		} else {
			panic("Wrong line: " + line)
		}
	}
	return ret
}

func atoi(s string) int {
	a, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return a
}

func inverseModule(a, n int) int {
	x, _, d := util.ExtendedGCD(a, n)
	if d != 1 {
		panic("Numbers are no coprime")
	}
	r := (n + x%n) % n
	if util.SafeMultModulo(r, a, n) != 1 {
		panic("Something went wrong")
	}
	return r
}

type Gen2 struct {
	N int
	V util.Matrix
}

func NewGen2(N int) Gen2 {
	g := Gen2{
		N: N,
		V: util.NewMatrix(3, 1),
	}
	*g.a() = 1
	*g.c() = 1
	return g
}

func (g Gen2) a() *int {
	return &g.V.M[0][0]
}

func (g Gen2) b() *int {
	return &g.V.M[1][0]
}

func (g Gen2) c() *int {
	return &g.V.M[2][0]
}

func (g Gen2) Get(i int) int {
	return normalize(util.SafeMultModulo(*g.a(), i, g.N)+*g.b(), g.N)
}

func (g Gen2) Print(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", g.Get(i))
	}
	fmt.Println()
}

type Transformation struct {
	N int
	M util.Matrix
}

func NewTransformation(N int, M [][]int) Transformation {
	t := Transformation{
		N: N,
		M: util.NewMatrix(3, 3),
	}
	for r := range t.M.M {
		for c := range t.M.M[r] {
			t.M.M[r][c] = normalize(M[r][c], t.N)
		}
	}
	return t
}

func IdentityTransformation(N int) Transformation {
	return NewTransformation(N, [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
}

func normalize(v, N int) int {
	return (N + v%N) % N
}

func normalizeMatrix(m util.Matrix, N int) util.Matrix {
	for r := range m.M {
		for c := range m.M[r] {
			m.M[r][c] = normalize(m.M[r][c], N)
		}
	}
	return m
}

func (t Transformation) Apply(g Gen2) Gen2 {
	v, err := t.M.Multiply(g.V, g.N)
	if err != nil {
		panic(err)
	}
	return Gen2{
		N: g.N,
		V: normalizeMatrix(v, g.N),
	}
}

func (t Transformation) Compose(t1 Transformation) Transformation {
	M, err := t1.M.Multiply(t.M, t.N)
	if err != nil {
		panic(err)
	}
	return Transformation{
		N: t.N,
		M: normalizeMatrix(M, t.N),
	}
}
func dealIntoNewStack(N int) Transformation {
	// i -> f(i) = a*i + b
	// 0 -> f(0)
	// n-1 -> f(n-1)
	// 0 -> f(n-1)
	// i -> f(n-1-i)
	// i -> a*(N-1-i) + b
	// i -> a*(N-1) -a*i +b
	// i -> -a*i + (a*(N-1) + b
	// i -> -a*i + (b-a)
	return NewTransformation(N, [][]int{
		{-1, 0, 0},
		{-1, 1, 0},
		{0, 0, 1},
	})
}

func cutNCards(N, n int) Transformation {
	// i -> f(i)
	// i -> f(i+n)
	// i -> a*i +a*n+b
	return NewTransformation(N, [][]int{
		{1, 0, 0},
		{n, 1, 0},
		{0, 0, 1},
	})
}

func dealWithIncrementN(N, n int) Transformation {
	// 0 -> f(0)
	// n*i -> f(i)
	// i=n*i
	// i -> f(i * 1/n)
	// i -> a*1/n * i + b
	return NewTransformation(N, [][]int{
		{inverseModule(n, N), 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
}

func test2() {
	N := 10

	g := NewGen2(N)
	g.Print(N)
	dealIntoNewStack(N).Apply(g).Print(N)

	fmt.Println()
	g = NewGen2(N)
	*g.b() = 2
	g.Print(N)
	dealIntoNewStack(N).Apply(g).Print(N)

	fmt.Println()
	g = NewGen2(N)
	*g.a() = 3
	g.Print(N)
	dealIntoNewStack(N).Apply(g).Print(N)

	fmt.Println()
	g = NewGen2(N)
	*g.a() = 3
	*g.b() = 2
	g.Print(N)
	g = dealIntoNewStack(N).Apply(g)
	g.Print(N)
	dealIntoNewStack(N).Apply(g).Print(N)

	fmt.Println()
	g = NewGen2(N)
	g.Print(N)
	cutNCards(N, 3).Apply(g).Print(N)

	fmt.Println()
	g = NewGen2(N)
	g.Print(N)
	cutNCards(N, -4).Apply(g).Print(N)

	fmt.Println()
	g = NewGen2(N)
	g.Print(N)
	dealWithIncrementN(N, 3).Apply(g).Print(N)
}

func test3() {
	N := 10

	g := NewGen2(N)

	t1 := dealIntoNewStack(N)
	t2 := cutNCards(N, -4)
	t3 := dealWithIncrementN(N, 3)

	T := IdentityTransformation(N)
	T = T.Compose(t1)
	T = T.Compose(t2)
	T = T.Compose(t3)

	T1 := T
	T1 = T1.Compose(T1)
	T2 := T.Compose(t1).Compose(t2).Compose(t3)

	g.Print(N)
	T1.Apply(g).Print(N)
	T2.Apply(g).Print(N)
	t3.Apply(t2.Apply(t1.Apply(t3.Apply(t2.Apply(t1.Apply(g)))))).Print(N)
}
