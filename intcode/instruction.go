package intcode

type instruction struct {
	code      int
	numParams int
	execute   func(c Computer, params []int)
}
