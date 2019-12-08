package intcode

type instruction struct {
	code         int
	numParams    int
	execute      func(c instructionContext, params []int) error
	modeOverride map[int]int
}

type instructionContext interface {
	stop()
	input() int
	output(int)
	setInstructionPointer(int)
	set(i, val int)
}

func (i *instruction) overrideMode(iParam, mode int) int {
	if m, ok := i.modeOverride[iParam]; ok {
		return m
	}
	return mode
}
