package day10

import (
	"strconv"
	"strings"
)

type CommSystem struct {
	x      int
	pc     int
	cycles []int
}

func LoadProgram(cmds []string) CommSystem {
	cs := CommSystem{x: 1, pc: 0, cycles: []int{}}

	for _, cmd := range cmds {
		switch cmd {
		case "noop":
			cs.cycles = append(cs.cycles, cs.x)
		default:
			dx, _ := strconv.Atoi(strings.Split(cmd, " ")[1])
			cs.cycles = append(cs.cycles, cs.x)
			cs.pc++
			cs.cycles = append(cs.cycles, cs.x)
			cs.x += dx
		}
		cs.pc++
	}

	return cs
}

func (cs CommSystem) Cycles() []int {
	return cs.cycles
}

func (cs CommSystem) RegisterX() int {
	return cs.x
}

func (cs CommSystem) SignalStrengthCycleN(n int) int {
	return cs.cycles[n-1] * n
}
