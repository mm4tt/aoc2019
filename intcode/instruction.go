package intcode

type instruction struct {
	code         int
	numParams    int
	execute      func(c Computer, params []int) error
	modeOverride map[int]int
}

func (i *instruction) overrideMode(iParam, mode int) int {
	if m, ok := i.modeOverride[iParam]; ok {
		return m
	}
	return mode
}
