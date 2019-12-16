package main

import (
	"fmt"

	"github.com/mm4tt/aoc2019/util"
)

func main() {
	nPhases, input := 100, "59704438946400225486037825889922820489843190285276623851650874501661128988396696069718826434708024511422795921838800269789913960190601300910423350290846455187315936154437526204822336114717910853157866334743979157700934791877134865819338701289349073169567308015162696370931073040617799608862983736292169088603858502137085782889297989277130087242942506416164598910622349994697403064628500493847458293153920207889114082230150603182206031692080645433361960358161328125435922180533297727179785114625861941781083443388701883640778753411135944703959349861504604264349715262460922987816868400261327556306957183739232107401756998929158348201149705670138765039"
	//nPhases, input = 4, "12345678"
	//nPhases, input = 100, "80871224585914546619083218645595"
	//nPhases, input = 100, "19617804207202209144916044189917"
	//nPhases, input = 100, "69317163492948606335995924319873"
	//nPhases, input = 100, "03036732577212944063491565474664"
	//nPhases, input = 100, "02935109699940807407585447034323"
	//nPhases, input = 100, "03081770884921959731165446850517"

	signal := make([]int, len(input))
	for i, c := range input {
		signal[i] = int(c) - int('0')
	}

	fmt.Println(solution1(signal, nPhases)[:8])
	fmt.Println()

	// Part2

	signal2 := make([]int, len(signal)*10000)
	for i := range signal2 {
		signal2[i] = signal[i%len(signal)]
	}
	offset := 0
	for i := 0; i < 7; i++ {
		offset = 10*offset + signal[i]
	}
	fmt.Println("Offset: ", offset)
	output := solution2(signal2, nPhases)
	fmt.Println(output[offset : offset+8])
}

func solution1(signal []int, nPhases int) []int {
	signal = append(signal[:0:0], signal...)
	n := len(signal)
	newSignal := make([]int, n)
	for ; nPhases > 0; nPhases-- {
		for i, _ := range signal {
			newSignal[i] = 0
			for j, v := range signal {
				newSignal[i] = newSignal[i] + getPatternI(i+1, j)*v
			}
			newSignal[i] = util.Abs(newSignal[i] % 10)
		}
		tmp := signal
		signal = newSignal
		newSignal = tmp
	}
	return signal
}

func solution2(signal []int, nPhases int) []int {
	signal = append(signal[:0:0], signal...)
	n := len(signal)
	newSignal := make([]int, n)
	sum := make([]int, n)
	for iPhase := 0; iPhase < nPhases; iPhase++ {
		fmt.Println("Done", iPhase, "phases")
		for i, v := range signal {
			if i == 0 {
				sum[i] = v
			} else {
				sum[i] = sum[i-1] + v
			}
		}
		getSum := func(a, b int) int {
			if a >= n {
				return 0
			}
			c := 0
			if a > 0 {
				c = sum[a-1]
			}
			return sum[util.Min(b, n)-1] - c
		}

		for i, _ := range signal {
			newSignal[i] = 0
			for j := i; j < n; j += 4 * (i + 1) {
				newSignal[i] += getSum(j, j+(i+1))
				newSignal[i] -= getSum(j+2*(i+1), j+3*(i+1))
			}
			newSignal[i] = util.Abs(newSignal[i] % 10)

		}
		tmp := signal
		signal = newSignal
		newSignal = tmp
	}
	return signal
}

func getPatternI(pos, i int) int {
	i -= pos - 1
	if i < 0 {
		return 0
	}
	i %= 4 * pos
	i /= pos
	switch i {
	case 0:
		return 1
	case 2:
		return -1
	}
	return 0
}
