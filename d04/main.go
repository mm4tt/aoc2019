package main

import "fmt"

func main() {
	min, max := 178416, 676461
	isOk := func(a int) bool {
		return true &&
			isSixDigitLong(a) &&
			isNonDecreasing(a) &&
			hasTwoSameAdjacent(a)
	}
	isOk2 := func(a int) bool {
		return true &&
			isSixDigitLong(a) &&
			isNonDecreasing(a) &&
			hasTwoSameAdjacent2(a)
	}

	c1, c2 := 0, 0
	for i := min; i <= max; i++ {
		if isOk(i) {
			c1++
		}
		if isOk2(i) {
			c2++
		}
	}
	fmt.Println(c1, c2)
}

func isSixDigitLong(a int) bool {
	return a >= 100000 && a < 1000000
}

func isNonDecreasing(a int) bool {
	c, stop := processDigits(a)
	defer close(stop)

	prev := -1
	for d := range c {
		if prev != -1 && prev > d {
			return false
		}
		prev = d
	}
	return true
}

func hasTwoSameAdjacent(a int) bool {
	c, stop := processDigits(a)
	defer close(stop)

	prev := -1
	for d := range c {
		if prev == d {
			return true
		}
		prev = d
	}
	return false
}

func hasTwoSameAdjacent2(a int) bool {
	c, stop := processDigits(a)
	defer close(stop)

	current, count := -1, 0
	for d := range c {
		if current == d {
			count++
			continue
		}
		if count == 2 {
			return true
		}
		current = d
		count = 1
	}
	if count == 2 {
		return true
	}
	return false
}

func processDigits(a int) (out <-chan int, stop chan<- struct{}) {
	outCh, stopCh := make(chan int), make(chan struct{})
	go processDigitsInternal(a, outCh, stopCh)
	return outCh, stopCh
}

func processDigitsInternal(a int, out chan<- int, stopCh <-chan struct{}) {
	p := 10
	for ; p < a; p *= 10 {
	}
	for p /= 10; p >= 1; p /= 10 {
		select {
		case out <- a / p:
		case <-stopCh:
			break
		}
		a = a % p
	}
	close(out)
}
