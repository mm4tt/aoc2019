package intcode

var instructions = []*instruction {add, mult, stop}

var add = &instruction{
	code:      1,
	numParams: 3,
	execute: func(c Computer, params []int) {
		c.Set(params[2], c.Get(params[0]) + c.Get(params[1]))
	},
}

var mult = &instruction{
	code:      2,
	numParams: 3,
	execute: func(c Computer, params []int) {
		c.Set(params[2], c.Get(params[0]) * c.Get(params[1]))
	},
}

var stop = &instruction{
	code:      99,
	numParams: 0,
	execute: func(c Computer, params []int) {
		c.stop()
	},
}