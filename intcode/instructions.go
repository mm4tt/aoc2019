package intcode

import "fmt"

var instructions = []*instruction{
	add,
	mult,
	input,
	output,
	jumpIfTrue,
	jumpIfFalse,
	lessThan,
	equals,
	stop,
}

var add = &instruction{
	code:      1,
	numParams: 3,
	execute: func(c Computer, params []int) error {
		c.Set(params[2], params[0]+params[1])
		return nil
	},
	modeOverride: map[int]int{2: direct}, // Set position is always direct.
}

var mult = &instruction{
	code:      2,
	numParams: 3,
	execute: func(c Computer, params []int) error {
		c.Set(params[2], params[0]*params[1])
		return nil
	},
	modeOverride: map[int]int{2: direct}, // Set position is always direct.
}

var input = &instruction{
	code:      3,
	numParams: 1,
	execute: func(c Computer, params []int) error {
		val, ok := <-c.input()
		if !ok {
			return fmt.Errorf("called input on empty input channel")
		}
		c.Set(params[0], val)
		return nil
	},
	modeOverride: map[int]int{0: direct},
}

var output = &instruction{
	code:      4,
	numParams: 1,
	execute: func(c Computer, params []int) error {
		c.output(params[0])
		return nil
	},
}

var jumpIfTrue = &instruction{
	code:      5,
	numParams: 2,
	execute: func(c Computer, params []int) error {
		if params[0] != 0 {
			c.setInstructionPointer(params[1])
		}
		return nil
	},
}

var jumpIfFalse = &instruction{
	code:      6,
	numParams: 2,
	execute: func(c Computer, params []int) error {
		if params[0] == 0 {
			c.setInstructionPointer(params[1])
		}
		return nil
	},
}

var lessThan = &instruction{
	code:      7,
	numParams: 3,
	execute: func(c Computer, params []int) error {
		val := 0
		if params[0] < params[1] {
			val = 1
		}
		c.Set(params[2], val)
		return nil
	},
	modeOverride: map[int]int{2: direct},
}

var equals = &instruction{
	code:      8,
	numParams: 3,
	execute: func(c Computer, params []int) error {
		val := 0
		if params[0] == params[1] {
			val = 1
		}
		c.Set(params[2], val)
		return nil
	},
	modeOverride: map[int]int{2: direct},
}

var stop = &instruction{
	code:      99,
	numParams: 0,
	execute: func(c Computer, params []int) error {
		c.stop()
		return nil
	},
}
