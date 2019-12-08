package util

import "sync"

type Barrier struct {
	wg sync.WaitGroup
}

func (b *Barrier) Run(f func()) {
	b.wg.Add(1)
	go func() {
		f()
		b.wg.Done()
	}()
}

func (b *Barrier) Wait() {
	b.wg.Wait()
}