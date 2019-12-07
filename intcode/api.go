package intcode

type Computer interface {
	Load(*Input)
	Set(i, val int)
	Get(i int) int

	Run() (*Output, error)
	RunAsync() *AsyncOutput

	stop()
	input() <-chan int
	output(int)
	setInstructionPointer(int)
}

type Input struct {
	Memory []int

	Inputs []int
	// InputCh has priority over Inputs, if InputChannel is set the Inputs field is ignored.
	InputCh chan int
}

type Output struct {
	Outputs []int
}

type AsyncOutput struct {
	OutputCh <-chan int
	ErrorCh  <-chan error
}
