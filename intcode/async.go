package intcode

type asyncOutput struct {
	errorCh  <-chan error
	outputCh <-chan int
}

func (a *asyncOutput) Finalize() (*Output, error) {
	output := &Output{}
loop:
	for ; ; {
		select {
		case o, ok := <-a.outputCh:
			if !ok {
				break loop
			}
			output.Outputs = append(output.Outputs, o)
		case err := <-a.errorCh:
			return nil, err
		}
	}
	return output, nil
}

func (a *asyncOutput) Process(f func(o int, err error)) {
	for ; ; {
		select {
		case o, ok := <-a.outputCh:
			if !ok {
				return
			}
			f(o, nil)
		case err := <-a.errorCh:
			f(0, err)
			return
		}
	}
}
