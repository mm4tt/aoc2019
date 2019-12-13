package intcode

import (
	"github.com/go-errors/errors"
	"github.com/golang-collections/collections/queue"
	"sync"
)

func NewComputer() Computer {
	instrMap := make(map[int]*instruction)
	for _, instr := range instructions {
		instrMap[instr.code] = instr
	}
	return &computer{
		instructions:   instrMap,
		inputQueueCond: sync.NewCond(new(sync.Mutex)),
	}
}

type computer struct {
	instructions map[int]*instruction

	memory       []int
	i            int
	relativeBase int
	stopCh       chan struct{}
	outputCh     chan int
	linkedTo     Computer

	// Ideally, we'd have a blocking queue...
	inputQueue     *queue.Queue
	inputQueueCond *sync.Cond
	inputPromptCh  chan struct{}
}

func (c *computer) LoadMemory(memory []int) {
	c.memory = append(memory[:0:0], memory...)
	c.i, c.relativeBase = 0, 0
	func() {
		c.inputQueueCond.L.Lock()
		defer c.inputQueueCond.L.Unlock()
		c.inputQueue = queue.New()
		c.inputPromptCh = make(chan struct{}, 1)
	}()
}

func (c *computer) Input(input ...int) {
	for _, i := range input {
		func(i int) {
			c.inputQueueCond.L.Lock()
			defer c.inputQueueCond.L.Unlock()
			c.inputQueue.Enqueue(i)
			if c.inputQueue.Len() == 1 {
				select {
				case <-c.inputPromptCh:
				default:
				}
				c.inputQueueCond.Signal()
			}
		}(i)
	}
}

func (c *computer) LinkTo(to Computer) {
	c.linkedTo = to
}

func (c *computer) Run() (*Output, error) {
	return c.RunAsync().Finalize()
}

func (c *computer) RunAsync() AsyncOutput {
	c.outputCh = make(chan int)
	errCh := make(chan error)
	output := &asyncOutput{outputCh: c.outputCh, errorCh: errCh, inputPromptCh: c.inputPromptCh}
	c.stopCh = make(chan struct{})
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
				break
			}
		}
		close(c.outputCh)
	}()
	return output
}

func (c *computer) Get(i int) int {
	return c.memory[i]
}

func (c *computer) Set(i, val int) {
	c.memory[i] = val
}

func (c *computer) stop() {
	close(c.stopCh)
}

func (c *computer) input() int {
	c.inputQueueCond.L.Lock()
	defer c.inputQueueCond.L.Unlock()
	if c.inputQueue.Len() == 0 {
		select {
		case c.inputPromptCh <- struct{}{}:
		default:
		}
		c.inputQueueCond.Wait()
	}
	return c.inputQueue.Dequeue().(int)
}

func (c *computer) output(o int) {
	c.outputCh <- o
	if c.linkedTo != nil {
		c.linkedTo.Input(o)
	}
}

func (c *computer) setInstructionPointer(i int) {
	c.i = i
}

func (c *computer) incRelativeBase(i int) {
	c.relativeBase += i
}

func (c *computer) computeParams(opCode int, instr *instruction) ([]*int, error) {
	code := opCode / 100
	params := make([]*int, instr.numParams)
	for j := 0; j < instr.numParams; j++ {
		mode := code % 10
		params[j] = &c.memory[c.i+1+j]
		switch mode {
		case direct:
		case indirect:
			ptr, err := c.memoryPtr(*params[j])
			if err != nil {
				return nil, err
			}
			params[j] = ptr
		case relative:
			ptr, err := c.memoryPtr(c.relativeBase + *params[j])
			if err != nil {
				return nil, err
			}
			params[j] = ptr
		default:
			return nil, errors.Errorf("invalid opCode: %d", opCode)
		}
		code /= 10
	}
	return params, nil
}

func (c *computer) memoryPtr(i int) (*int, error) {
	if i < 0 {
		return nil, errors.Errorf("negative memory index: %d", i)
	}
	if i >= len(c.memory) {
		c.memory = append(c.memory, make([]int, i-len(c.memory)+1)...)
	}
	return &c.memory[i], nil
}
