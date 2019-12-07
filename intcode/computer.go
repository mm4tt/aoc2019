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

	memory   []int
	i        int
	stopCh   chan struct{}
	inputCh  chan int
	outputCh chan int
}

func (c *computer) Load(input *Input) {
	c.memory = append(input.Memory[:0:0], input.Memory...)
	c.i = 0
	c.stopCh = make(chan struct{})

	if input.InputCh != nil {
		c.inputCh = input.InputCh
	} else {
		c.inputCh = make(chan int, len(input.Inputs))
		for _, i := range input.Inputs {
			c.inputCh <- i
		}
		close(c.inputCh)
	}
}

func (c *computer) Run() (*Output, error) {
	asyncOutput := c.RunAsync()
	output := &Output{}
loop:
	for ; ; {
		select {
		case o, ok := <-asyncOutput.OutputCh:
			if !ok {
				break loop
			}
			output.Outputs = append(output.Outputs, o)
		case err := <-asyncOutput.ErrorCh:
			return nil, err
		}
	}

	return output, nil
}

func (c *computer) RunAsync() *AsyncOutput {
	c.outputCh = make(chan int)
	errCh := make(chan error)
	output := &AsyncOutput{OutputCh: c.outputCh, ErrorCh: errCh}
	go func() {
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
				errCh <- errors.Errorf("invalid opCode: %d", opCode)
				return
			}
			params, err := c.computeParams(opCode, instr)
			if err != nil {
				errCh <- err
				return
			}
			c.i += 1 + instr.numParams
			if err := instr.execute(c, params); err != nil {
				errCh <- err
				return
			}
		}
		close(c.outputCh)
	}()
	return output
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
	c.outputCh <- o
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
