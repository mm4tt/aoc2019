package main

import (
	"fmt"
	"github.com/mm4tt/aoc2019/intcode"
	"github.com/mm4tt/aoc2019/util"
	"log"
)

func main() {
	inputMem := []int{1, 330, 331, 332, 109, 3032, 1102, 1, 1182, 16, 1101, 0, 1433, 24, 102, 1, 0, 570, 1006, 570, 36, 1001, 571, 0, 0, 1001, 570, -1, 570, 1001, 24, 1, 24, 1106, 0, 18, 1008, 571, 0, 571, 1001, 16, 1, 16, 1008, 16, 1433, 570, 1006, 570, 14, 21102, 1, 58, 0, 1105, 1, 786, 1006, 332, 62, 99, 21101, 333, 0, 1, 21101, 73, 0, 0, 1106, 0, 579, 1101, 0, 0, 572, 1102, 0, 1, 573, 3, 574, 101, 1, 573, 573, 1007, 574, 65, 570, 1005, 570, 151, 107, 67, 574, 570, 1005, 570, 151, 1001, 574, -64, 574, 1002, 574, -1, 574, 1001, 572, 1, 572, 1007, 572, 11, 570, 1006, 570, 165, 101, 1182, 572, 127, 101, 0, 574, 0, 3, 574, 101, 1, 573, 573, 1008, 574, 10, 570, 1005, 570, 189, 1008, 574, 44, 570, 1006, 570, 158, 1105, 1, 81, 21102, 1, 340, 1, 1106, 0, 177, 21102, 1, 477, 1, 1105, 1, 177, 21102, 514, 1, 1, 21101, 0, 176, 0, 1106, 0, 579, 99, 21101, 184, 0, 0, 1105, 1, 579, 4, 574, 104, 10, 99, 1007, 573, 22, 570, 1006, 570, 165, 102, 1, 572, 1182, 21101, 375, 0, 1, 21101, 0, 211, 0, 1105, 1, 579, 21101, 1182, 11, 1, 21102, 222, 1, 0, 1105, 1, 979, 21101, 388, 0, 1, 21101, 233, 0, 0, 1105, 1, 579, 21101, 1182, 22, 1, 21101, 244, 0, 0, 1105, 1, 979, 21101, 0, 401, 1, 21102, 255, 1, 0, 1105, 1, 579, 21101, 1182, 33, 1, 21102, 1, 266, 0, 1106, 0, 979, 21102, 1, 414, 1, 21102, 1, 277, 0, 1106, 0, 579, 3, 575, 1008, 575, 89, 570, 1008, 575, 121, 575, 1, 575, 570, 575, 3, 574, 1008, 574, 10, 570, 1006, 570, 291, 104, 10, 21101, 0, 1182, 1, 21102, 1, 313, 0, 1105, 1, 622, 1005, 575, 327, 1102, 1, 1, 575, 21102, 1, 327, 0, 1105, 1, 786, 4, 438, 99, 0, 1, 1, 6, 77, 97, 105, 110, 58, 10, 33, 10, 69, 120, 112, 101, 99, 116, 101, 100, 32, 102, 117, 110, 99, 116, 105, 111, 110, 32, 110, 97, 109, 101, 32, 98, 117, 116, 32, 103, 111, 116, 58, 32, 0, 12, 70, 117, 110, 99, 116, 105, 111, 110, 32, 65, 58, 10, 12, 70, 117, 110, 99, 116, 105, 111, 110, 32, 66, 58, 10, 12, 70, 117, 110, 99, 116, 105, 111, 110, 32, 67, 58, 10, 23, 67, 111, 110, 116, 105, 110, 117, 111, 117, 115, 32, 118, 105, 100, 101, 111, 32, 102, 101, 101, 100, 63, 10, 0, 37, 10, 69, 120, 112, 101, 99, 116, 101, 100, 32, 82, 44, 32, 76, 44, 32, 111, 114, 32, 100, 105, 115, 116, 97, 110, 99, 101, 32, 98, 117, 116, 32, 103, 111, 116, 58, 32, 36, 10, 69, 120, 112, 101, 99, 116, 101, 100, 32, 99, 111, 109, 109, 97, 32, 111, 114, 32, 110, 101, 119, 108, 105, 110, 101, 32, 98, 117, 116, 32, 103, 111, 116, 58, 32, 43, 10, 68, 101, 102, 105, 110, 105, 116, 105, 111, 110, 115, 32, 109, 97, 121, 32, 98, 101, 32, 97, 116, 32, 109, 111, 115, 116, 32, 50, 48, 32, 99, 104, 97, 114, 97, 99, 116, 101, 114, 115, 33, 10, 94, 62, 118, 60, 0, 1, 0, -1, -1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 18, 14, 0, 109, 4, 2102, 1, -3, 587, 20101, 0, 0, -1, 22101, 1, -3, -3, 21102, 0, 1, -2, 2208, -2, -1, 570, 1005, 570, 617, 2201, -3, -2, 609, 4, 0, 21201, -2, 1, -2, 1106, 0, 597, 109, -4, 2106, 0, 0, 109, 5, 1201, -4, 0, 630, 20102, 1, 0, -2, 22101, 1, -4, -4, 21101, 0, 0, -3, 2208, -3, -2, 570, 1005, 570, 781, 2201, -4, -3, 652, 21001, 0, 0, -1, 1208, -1, -4, 570, 1005, 570, 709, 1208, -1, -5, 570, 1005, 570, 734, 1207, -1, 0, 570, 1005, 570, 759, 1206, -1, 774, 1001, 578, 562, 684, 1, 0, 576, 576, 1001, 578, 566, 692, 1, 0, 577, 577, 21101, 0, 702, 0, 1105, 1, 786, 21201, -1, -1, -1, 1106, 0, 676, 1001, 578, 1, 578, 1008, 578, 4, 570, 1006, 570, 724, 1001, 578, -4, 578, 21101, 0, 731, 0, 1106, 0, 786, 1106, 0, 774, 1001, 578, -1, 578, 1008, 578, -1, 570, 1006, 570, 749, 1001, 578, 4, 578, 21101, 0, 756, 0, 1105, 1, 786, 1105, 1, 774, 21202, -1, -11, 1, 22101, 1182, 1, 1, 21101, 774, 0, 0, 1105, 1, 622, 21201, -3, 1, -3, 1106, 0, 640, 109, -5, 2105, 1, 0, 109, 7, 1005, 575, 802, 20101, 0, 576, -6, 20101, 0, 577, -5, 1106, 0, 814, 21102, 0, 1, -1, 21102, 1, 0, -5, 21102, 0, 1, -6, 20208, -6, 576, -2, 208, -5, 577, 570, 22002, 570, -2, -2, 21202, -5, 41, -3, 22201, -6, -3, -3, 22101, 1433, -3, -3, 2102, 1, -3, 843, 1005, 0, 863, 21202, -2, 42, -4, 22101, 46, -4, -4, 1206, -2, 924, 21102, 1, 1, -1, 1105, 1, 924, 1205, -2, 873, 21102, 1, 35, -4, 1105, 1, 924, 1202, -3, 1, 878, 1008, 0, 1, 570, 1006, 570, 916, 1001, 374, 1, 374, 1201, -3, 0, 895, 1102, 2, 1, 0, 1202, -3, 1, 902, 1001, 438, 0, 438, 2202, -6, -5, 570, 1, 570, 374, 570, 1, 570, 438, 438, 1001, 578, 558, 922, 20101, 0, 0, -4, 1006, 575, 959, 204, -4, 22101, 1, -6, -6, 1208, -6, 41, 570, 1006, 570, 814, 104, 10, 22101, 1, -5, -5, 1208, -5, 39, 570, 1006, 570, 810, 104, 10, 1206, -1, 974, 99, 1206, -1, 974, 1101, 0, 1, 575, 21101, 973, 0, 0, 1106, 0, 786, 99, 109, -7, 2105, 1, 0, 109, 6, 21101, 0, 0, -4, 21102, 1, 0, -3, 203, -2, 22101, 1, -3, -3, 21208, -2, 82, -1, 1205, -1, 1030, 21208, -2, 76, -1, 1205, -1, 1037, 21207, -2, 48, -1, 1205, -1, 1124, 22107, 57, -2, -1, 1205, -1, 1124, 21201, -2, -48, -2, 1105, 1, 1041, 21101, 0, -4, -2, 1105, 1, 1041, 21101, -5, 0, -2, 21201, -4, 1, -4, 21207, -4, 11, -1, 1206, -1, 1138, 2201, -5, -4, 1059, 2101, 0, -2, 0, 203, -2, 22101, 1, -3, -3, 21207, -2, 48, -1, 1205, -1, 1107, 22107, 57, -2, -1, 1205, -1, 1107, 21201, -2, -48, -2, 2201, -5, -4, 1090, 20102, 10, 0, -1, 22201, -2, -1, -2, 2201, -5, -4, 1103, 2101, 0, -2, 0, 1106, 0, 1060, 21208, -2, 10, -1, 1205, -1, 1162, 21208, -2, 44, -1, 1206, -1, 1131, 1105, 1, 989, 21102, 1, 439, 1, 1106, 0, 1150, 21102, 1, 477, 1, 1105, 1, 1150, 21101, 514, 0, 1, 21101, 0, 1149, 0, 1106, 0, 579, 99, 21102, 1, 1157, 0, 1105, 1, 579, 204, -2, 104, 10, 99, 21207, -3, 22, -1, 1206, -1, 1138, 2101, 0, -5, 1176, 2101, 0, -4, 0, 109, -6, 2105, 1, 0, 8, 7, 34, 1, 5, 1, 34, 1, 5, 1, 34, 1, 5, 1, 32, 11, 30, 1, 1, 1, 5, 1, 1, 1, 30, 1, 1, 1, 5, 7, 26, 1, 1, 1, 7, 1, 3, 1, 26, 1, 1, 1, 5, 11, 22, 1, 1, 1, 5, 1, 1, 1, 3, 1, 3, 1, 16, 7, 1, 9, 3, 1, 3, 1, 5, 1, 10, 1, 13, 1, 5, 1, 3, 1, 5, 1, 10, 1, 13, 1, 5, 1, 3, 1, 5, 1, 10, 1, 13, 1, 5, 1, 3, 1, 5, 1, 10, 1, 13, 1, 3, 7, 5, 1, 10, 1, 13, 1, 5, 1, 9, 1, 10, 1, 13, 7, 9, 12, 39, 12, 29, 1, 10, 1, 29, 1, 10, 1, 29, 1, 10, 1, 29, 1, 8, 7, 25, 1, 8, 1, 1, 1, 3, 1, 25, 1, 4, 7, 3, 1, 19, 7, 4, 1, 3, 1, 5, 1, 19, 1, 10, 1, 3, 1, 5, 1, 19, 1, 10, 1, 3, 1, 5, 1, 19, 1, 10, 1, 3, 1, 5, 1, 7, 7, 5, 1, 10, 1, 3, 1, 5, 1, 7, 1, 5, 1, 5, 1, 10, 11, 7, 1, 5, 1, 5, 1, 14, 1, 13, 1, 5, 1, 5, 1, 14, 7, 7, 1, 5, 1, 5, 1, 20, 1, 7, 1, 5, 1, 5, 1, 20, 1, 7, 1, 5, 7, 20, 1, 7, 1, 32, 1, 7, 1, 32, 1, 7, 1, 32, 9, 18}

	c := intcode.NewComputer()
	c.LoadMemory(inputMem)

	output, err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
	board := []string{}
	line := []rune{}
	for _, i := range output.Outputs {
		c := rune(i)
		if c == '\n' {
			board = append(board, string(line))
			line = []rune{}
			continue
		}
		line = append(line, c)
	}

	isIntersection := func(x, y int) bool {
		return board[y][x] == '#' &&
			board[y+1][x] == '#' &&
			board[y-1][x] == '#' &&
			board[y][x+1] == '#' &&
			board[y][x-1] == '#'

	}

	draw(board)
	Y, X := len(board)-1, len(board[0])

	alignments := 0
	var p util.Point2D
	for y := 1; y+1 < Y; y++ {
		for x := 1; x+1 < X; x++ {
			//fmt.Println(y, x)
			if isIntersection(x, y) {
				alignments += x * y
			}
			if board[y][x] == '^' {
				p = util.Point2D{x, y}
			}
		}
	}
	fmt.Println(alignments)

	// Part 2
	dir := util.Point2D{0, -1}
	instr := ""

	isOnPath := func(p util.Point2D) bool {
		return true &&
			0 <= p.X && p.X < X &&
			0 <= p.Y && p.Y < Y &&
			board[p.Y][p.X] == '#'
	}

	currentMove := 0
	for ; ; {
		n := p.Add(dir)
		if isOnPath(n) {
			currentMove++
			p = n
			continue
		}
		if currentMove > 0 {
			instr += fmt.Sprintf("%d,", currentMove)
			currentMove = 0
		}
		if isOnPath(p.Add(dir.RotateLeft())) {
			instr += "L,"
			dir = dir.RotateLeft()
			continue
		}
		if isOnPath(p.Add(dir.RotateRight())) {
			instr += "R,"
			dir = dir.RotateRight()
			continue
		}
		break
	}
	fmt.Println(instr)

	inputStr := "A,B,A,B,C,A,B,C,A,C\nR,6,L,6,L,10\nL,8,L,6,L,10,L,6\nR,6,L,8,L,10,R,6\nn\n"
	input := make([]int, len(inputStr))
	for i, c := range inputStr {
		input[i] = int(c)
	}
	c.LoadMemory(inputMem)
	c.Set(0, 2)
	c.Input(input...)
	output, err = c.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output.Outputs[len(output.Outputs)-1])
}

func draw(board []string) {
	for _, line := range board {
		fmt.Println(line)
	}
}