package main

import "fmt"

type Gen struct {
	a, b, N int
	// i -> (a*i + b) % N
}

func newGen(N int) Gen {
	return Gen{
		N: N,
		a: 1,
		b: 0,
	}
}

func (g Gen) normalize(v int) int {
	return (g.N + v%g.N) % g.N
}

func (g Gen) get(i int) int {
	return (g.a*i + g.b) % g.N
}

func (g Gen) dealIntoNewStack() Gen {
	// i -> a*i + b
	// 0 -> f(0)
	// n-1 -> f(n-1)
	// 0 -> f(n-1)
	// i -> f(n-1-i)
	// i -> a*(N-1-i) + b
	// i -> a*(N-1) -a*i +b
	// i -> -a*i + (a*(N-1) + b
	return Gen{
		N: g.N,
		a: g.normalize(-g.a),
		b: g.normalize(g.b - g.a),
	}
}

func (g Gen) cutNCards(n int) Gen {
	// i -> f(i)
	// i -> f(i+n)
	// i -> a*i +a*n+b
	return Gen{
		N: g.N,
		a: g.a,
		b: g.normalize(g.a*n + g.b),
	}
}

func (g Gen) dealWithIncrementN(n int) Gen {
	// 0 -> f(0)
	// n*i -> f(i)
	// i=n*i
	// i -> f(i * 1/n)
	// i -> a*1/n * i + b
	return Gen{
		N: g.N,
		a: g.normalize(g.a * inverseModule(n, g.N)),
		b: g.b,
	}
}

func (g Gen) print(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", g.get(i))
	}
	fmt.Println()
}

func test() {
	g := newGen(10)
	g.print(10)
	g.dealIntoNewStack().print(10)

	fmt.Println()
	g = newGen(10)
	g.b = 2
	g.print(10)
	g.dealIntoNewStack().print(10)

	fmt.Println()
	g = newGen(10)
	g.a = 3
	g.print(10)
	g.dealIntoNewStack().print(10)

	fmt.Println()
	g = newGen(10)
	g.a = 3
	g.b = 2
	g.print(10)
	g.dealIntoNewStack().print(10)

	fmt.Println()
	g = newGen(10)
	g.print(10)
	g = g.cutNCards(3)
	g.print(10)

	fmt.Println()
	g = newGen(10)
	g.print(10)
	g = g.cutNCards(-4)
	g.print(10)

	fmt.Println()
	g = newGen(10)
	g.print(10)
	g = g.dealWithIncrementN(3)
	g.print(10)
}
