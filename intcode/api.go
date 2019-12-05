package intcode

type Computer interface {
	Load(*Input)
	Run() (*Output, error)
	Set(i, val int)
	Get(i int) int

	stop()
	input() <-chan int
	output(int)
	setInstructionPointer(int)
}

type Input struct {
	Memory []int
	Inputs []int
}

type Output struct {
	Outputs []int
}
