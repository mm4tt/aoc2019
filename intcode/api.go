package intcode

type Computer interface {
	LoadMemory([]int)
	Set(i, val int)
	Get(i int) int

	Input(...int)

	// LinkTo will connect output of this Computer to the input of the provided Computer.
	LinkTo(Computer)

	Run() (*Output, error)
	// RunAsync runs the program in the async mode. The Finalize() method must be always called on
	// the returned AsyncOutput before any other program is run on this computer and / or to avoid
	// goroutine leaks.
	RunAsync() AsyncOutput
}

type Output struct {
	Outputs []int
}

type AsyncOutput interface {
	Finalize() (*Output, error)
}
