package main

import (
	"fmt"
	"github.com/mm4tt/aoc2019/util"
	"log"
	"strconv"
	"strings"
)

func main() {
	nSteps := 1000
	lines, err := util.ReadLines("$GOPATH/src/github.com/mm4tt/aoc2019/d12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	startMoons, startVels := make([]util.Point3d, len(lines)), make([]util.Point3d, len(lines))
	for line := range lines {
		startMoons = append(startMoons, parsePoint(line))
		startVels = append(startVels, util.Point3d{0, 0, 0})
	}

	// Part 1
	func(moons, vels []util.Point3d) {
		for ; nSteps > 0; nSteps-- {
			for i, m := range moons {
				for j := i + 1; j < len(moons); j++ {
					g := calculateGravity(m, moons[j])
					vels[i] = vels[i].Add(g.Neg())
					vels[j] = vels[j].Add(g)
				}
			}
			for i, v := range vels {
				moons[i] = moons[i].Add(v)
			}
		}

		total := 0
		for i := range moons {
			total += moons[i].Norm() * vels[i].Norm()
		}
		fmt.Println(total)
	}(append(startMoons[:0:0], startMoons...), append(startVels[:0:0], startVels...))

	// Part 2
	func(moons, vels []util.Point3d) {
		xPeriod, yPeriod, zPeriod := 0, 0, 0
		for k := 1; xPeriod*yPeriod*zPeriod == 0; k++ {
			for i, m := range moons {
				for j := i + 1; j < len(moons); j++ {
					g := calculateGravity(m, moons[j])
					vels[i] = vels[i].Add(g.Neg())
					vels[j] = vels[j].Add(g)
				}
			}
			for i, v := range vels {
				moons[i] = moons[i].Add(v)
			}

			if xPeriod == 0 && allEqual(startMoons, moons, xEqual) && allEqual(startVels, vels, xEqual) {
				xPeriod = k
			}
			if yPeriod == 0 && allEqual(startMoons, moons, yEqual) && allEqual(startVels, vels, yEqual) {
				yPeriod = k
			}
			if zPeriod == 0 && allEqual(startMoons, moons, zEqual) && allEqual(startVels, vels, zEqual) {
				zPeriod = k
			}
		}

		fmt.Println(xPeriod, yPeriod, zPeriod)
		fmt.Println(util.LMC(xPeriod, yPeriod, zPeriod))

	}(append(startMoons[:0:0], startMoons...), append(startVels[:0:0], startVels...))

}

func parsePoint(line string) util.Point3d {
	line = strings.Trim(line, "<>\n")
	elems := strings.Split(line, ", ")
	f := func(s string) int {
		i, err := strconv.Atoi(s[2:])
		if err != nil {
			panic(err)
		}
		return i
	}
	return util.Point3d{f(elems[0]), f(elems[1]), f(elems[2])}
}

func calculateGravity(m1, m2 util.Point3d) util.Point3d {
	return util.Point3d{X: sgn(m1.X - m2.X), Y: sgn(m1.Y - m2.Y), Z: sgn(m1.Z - m2.Z),}
}

func sgn(a int) int {
	if a == 0 {
		return 0
	}
	return a / util.Abs(a)
}

func allEqual(a, b []util.Point3d, equal func(a, b util.Point3d) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func xEqual(a, b util.Point3d) bool {
	return a.X == b.X
}

func yEqual(a, b util.Point3d) bool {
	return a.Y == b.Y
}

func zEqual(a, b util.Point3d) bool {
	return a.Z == b.Z
}
