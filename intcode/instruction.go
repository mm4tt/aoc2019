package intcode

type instruction struct {
	code         int
	numParams    int
	execute      func(c instructionContext, params []*int) error
}

type instructionContext interface {
	stop()
	input() int
	output(int)
	setInstructionPointer(int)
}
