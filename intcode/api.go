package intcode

type Computer interface {
	Load(input Input)
	Run() (Output, error)
	Set(i, val int)
	Get(i int) int

	stop()
}

type Input struct {
	Memory []int
	Inputs []int
}

type Output struct {
	Output []int
}
