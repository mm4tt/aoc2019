package main

import (
	"fmt"
	"github.com/mm4tt/aoc2019/util"
	"log"
	"strconv"
	"strings"
)

type Chemical string
type QuantizedChemical struct {
	c Chemical
	q int
}
type Reaction struct {
	input  []QuantizedChemical
	output QuantizedChemical
}

const (
	FUEL = Chemical("FUEL")
	ORE  = Chemical("ORE")
)

func main() {
	lines, err := util.ReadLines("d14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	reactions := make(map[Chemical]Reaction)
	for line := range lines {
		c := strings.Split(line, "=>")
		var r Reaction
		qc, err := parseQuantizedChemical(c[1])
		if err != nil {
			log.Fatal(err)
		}
		r.output = qc
		for _, s := range strings.Split(c[0], ", ") {
			qc, err := parseQuantizedChemical(s)
			if err != nil {
				log.Fatal(err)
			}
			r.input = append(r.input, qc)
		}
		reactions[r.output.c] = r
	}

	// Part1
	func() {
		requirements, extra := make(map[Chemical]int), make(map[Chemical]int)
		requirements[FUEL] = 1
		func() {
			for ; ; {
				fmt.Println("Requirements: ", requirements)
				fmt.Println("Extra: ", extra)
				var c Chemical
				for c, _ = range requirements {
					if c == ORE {
						if len(requirements) == 1 {
							return // Nothing more to process
						}
					} else {
						break
					}
				}
				func() {
					defer delete(requirements, c)
					q := requirements[c]
					fmt.Println("Processing:", c, ", needing: ", q)
					if e := extra[c]; e > 0 {
						fmt.Println("\tHave extra ", e, "of", c)
						q1 := util.Max(q-e, 0)
						extra[c] = extra[c] - (q - q1)
						q = q1
						fmt.Println("\tAfter using from extra needing", q)
					}

					if q == 0 {
						fmt.Println("\tTook everything from extra, done")
						return
					}
					r := reactions[c]
					fmt.Println("\tReaction is: ", r)
					n := util.Ceil(q, r.output.q)
					fmt.Println("\tNeed to execute it", n, "times")
					for _, cq := range r.input {
						requirements[cq.c] = requirements[cq.c] + cq.q*n
					}
					extra[c] = extra[c] + n*r.output.q - q
				}()
			}
		}()
		fmt.Println("\nIt will require:", requirements)
	}()

	// Part2
	canProduce := func(nFuel int) bool {
		requirements, extra := make(map[Chemical]int), make(map[Chemical]int)
		requirements[FUEL] = nFuel
		func() {
			for ; len(requirements) > 0; {
				var c Chemical
				for c, _ = range requirements {
					if c == ORE {
						if len(requirements) == 1 {
							return // Nothing more to process
						}
					} else {
						break
					}
				}
				func() {
					defer delete(requirements, c)
					q := requirements[c]
					if e := extra[c]; e > 0 {
						q1 := util.Max(q-e, 0)
						extra[c] = extra[c] - (q - q1)
						q = q1
					}

					if q == 0 {
						return
					}
					r := reactions[c]
					n := util.Ceil(q, r.output.q)
					for _, cq := range r.input {
						requirements[cq.c] = requirements[cq.c] + cq.q*n
					}
					extra[c] = extra[c] + n*r.output.q - q
				}()
			}
		}()
		return requirements[ORE] <= 1000000000000
	}

	lastOk, i := 1, 2
	for ; canProduce(i); {
		lastOk, i = i, i*2
	}
	a, b := lastOk, i
	for ; b-a > 1; {
		i = (a + b) / 2
		if canProduce(i) {
			a = i
		} else {
			b = i
		}
	}
	fmt.Println("\nFuel that can be produced with 1000000000000 ORE: ", a)
}

func parseQuantizedChemical(s string) (QuantizedChemical, error) {
	s = strings.Trim(s, " ")
	a := strings.Split(s, " ")
	q, err := strconv.Atoi(a[0])
	if err != nil {
		return QuantizedChemical{}, err
	}
	return QuantizedChemical{c: Chemical(a[1]), q: q}, nil
}
