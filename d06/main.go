package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func main() {
	inputFile, err := os.Open(os.ExpandEnv("$GOPATH/src/github.com/mm4tt/aoc2019/d06/input.txt"))
	if err != nil {
		log.Fatal(err)
	}

	g, rg := make(map[string][]string), make(map[string]string)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		objects := strings.Split(line, ")")
		a, b := objects[0], objects[1]
		g[a] = append(g[a], b)
		rg[b] = a
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1.
	func() {
		type elem struct {
			node  string
			level int
		}
		s := stack.New()
		s.Push(elem{"COM", 0})
		checksum := 0
		for ; s.Len() > 0; {
			e := s.Pop().(elem)
			checksum += e.level
			for _, n := range g[e.node] {
				s.Push(elem{n, e.level + 1})
			}
		}
		fmt.Println(checksum)
	}()

	// Part 2.
	func() {
		s := make(map[string]int)
		n := "YOU"
		for i := 0; rg[n] != ""; i++ {
			n = rg[n]
			s[n] = i
		}

		n = "SAN"
		for i := 0; rg[n] != ""; i++ {
			n = rg[n]
			if j, ok := s[n]; ok {
				fmt.Println(i + j)
			}
		}
		panic("no common ancestor...")
	}()

}
