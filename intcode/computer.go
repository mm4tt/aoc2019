package intcode

import (
	"github.com/go-errors/errors"
)

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

	memory  []int
	i       int
	stopCh  chan struct{}
	inputCh chan int
	outputs []int
}

func (c *computer) Load(input *Input) {
	c.memory = append(input.Memory[:0:0], input.Memory...)
	c.i = 0
	c.stopCh = make(chan struct{})
	c.inputCh = make(chan int, len(input.Inputs))
	for _, i := range input.Inputs {
		c.inputCh <- i
	}
}

func (c *computer) Run() (*Output, error) {
	c.outputs = []int{}
loop:
	for ; ; {
		select {
		case <-c.stopCh:
			break loop
		default:
		}

		opCode := c.memory[c.i]
		instr, ok := c.instructions[opCode%100]
		if !ok {
			return nil, errors.Errorf("invalid opCode: %d", opCode)
		}
		params, err := c.computeParams(opCode, instr)
		if err != nil {
			return nil, err
		}
		c.i += 1 + instr.numParams
		if err := instr.execute(c, params); err != nil {
			return nil, err
		}

	}
	return &Output{Outputs: c.outputs}, nil
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

func (c *computer) input() <-chan int {
	return c.inputCh
}

func (c *computer) output(o int) {
	c.outputs = append(c.outputs, o)
}

func (c *computer) setInstructionPointer(i int) {
	c.i = i
}


func (c *computer) computeParams(opCode int, instr *instruction) ([]int, error) {
	code := opCode / 100
	params := make([]int, instr.numParams)
	for j := 0; j < instr.numParams; j++ {
		mode := instr.overrideMode(j, code%10)
		params[j] = c.memory[c.i+1+j]
		switch mode {
		case indirect:
			params[j] = c.memory[params[j]]
		case direct:
		default:
			return nil, errors.Errorf("invalid opCode: %d", opCode)
		}
		code /= 10
	}
	return params, nil
}
