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
