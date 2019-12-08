package main

import (
	"fmt"
	"github.com/mm4tt/aoc2019/util"
	"log"
)

func main() {
	input, err := util.ReadString("$GOPATH/src/github.com/mm4tt/aoc2019/d08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	w, h := 25, 6

	//input, w, h = "123456789012", 3, 2

	n := w * h

	// Part 1
	func() {
		var bestLayerCounts map[rune]int
		for i := 0; i < len(input); i += n {
			layer := input[i : i+n]
			counts := getLayerCounts(layer)
			if bestLayerCounts == nil || bestLayerCounts['0'] > counts['0'] {
				bestLayerCounts = counts
			}
		}
		fmt.Println(bestLayerCounts['1'] * bestLayerCounts['2'])
	}()

	// Part 2
	func() {
		image := make([]rune, n)
		for i := range image {
			image[i] = '2'
		}
		for i := 0; i < len(input); i += n {
			layer := input[i : i+n]
			for j, c := range layer {
				if image[j] == '2' {
					image[j] = c
				}
			}
		}

		for i := range image {
			if image[i] == '1' {
				image[i] = '#'
			} else {
				image[i] = ' '
			}
		}
		for i := 0; i < len(image); i += w {
			fmt.Println(string(image[i : i+w]))
		}
	}()
}

func getLayerCounts(layer string) map[rune]int {
	m := make(map[rune]int)
	for _, c := range layer {
		m[c] = m[c] + 1
	}
	return m
}
