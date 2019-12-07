package util

func Permutate(a []int) <-chan[]int {
	output := make(chan []int)
	go func() {
		permutate(a, []int {}, output)
		close(output)
	}()
	return output
}

func permutate(items, current []int, output chan<-[]int) {
	if len(items) == 0 {
		output <- current
		return
	}

	for i, e := range items {
		permutate(append(append(items[:0:0], items[:i]...), items[i+1:]...), append(current, e), output)
	}
}
