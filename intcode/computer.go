package intcode

import "fmt"

func NewComputer() Computer {
	instrMap := make(map[int]*instruction)
	for _, instr := range instructions {
		instrMap[instr.code] = instr
	}
	return &computer{
		instructions: instrMap,
	}
}

type computer struct {
	instructions map[int]*instruction

	memory []int
	i      int
	stopCh chan struct{}
}

func (c *computer) Load(input Input) {
	c.memory = append(input.Memory[:0:0], input.Memory...)
	c.i = 0
	c.stopCh = make(chan struct{})
	// TODO: inputs
}

func (c *computer) Run() (Output, error) {
	output := Output{}

loop:
	for ; ; {
		select {
		case <-c.stopCh:
			break loop
		default:
		}

		opCode := c.memory[c.i]
		instr, ok := c.instructions[opCode]
		if !ok {
			return output, fmt.Errorf("invalid opCode: %d", opCode)
		}
		instr.execute(c, c.memory[c.i+1:c.i+1+instr.numParams])
		c.i += 1 + instr.numParams
	}

	return output, nil
}

func (c *computer) Set(i, val int) {
	c.memory[i] = val
}

func (c *computer) Get(i int) int {
	return c.memory[i]
}

func (c *computer) stop() {
	close(c.stopCh)
}
