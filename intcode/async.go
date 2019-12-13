package intcode

type asyncOutput struct {
	errorCh       <-chan error
	outputCh      <-chan int
	inputPromptCh <-chan struct{}
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

func (a *asyncOutput) Process(f func(AsyncProcessEvent)) {
	for ; ; {
		select {
		case o, ok := <-a.outputCh:
			if !ok {
				return
			}
			f(AsyncProcessEvent{o, nil, false})
		case err := <-a.errorCh:
			f(AsyncProcessEvent{0, err, false})
			return
		case <-a.inputPromptCh:
			f(AsyncProcessEvent{0, nil, true})
		}
	}
}
