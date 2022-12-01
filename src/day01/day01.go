package day1

import (
	"sort"
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/mathutil"
)

type Day1 struct {
	Input string
	elves [][]int
}

func (d Day1) MaxCaloriesSingleElf() int {
	if d.elves == nil {
		d.init()
	}

	maxCalories := 0
	for _, elf := range d.elves {
		maxCalories = mathutil.Max(maxCalories, mathutil.Sum(elf...))
	}

	return maxCalories
}

func (d Day1) CaloriesTopThreeElves() int {
	if d.elves == nil {
		d.init()
	}

	elfCalories := []int{}
	for _, elf := range d.elves {
		elfCalories = append(elfCalories, mathutil.Sum(elf...))
	}

	sort.Ints(elfCalories)

	return mathutil.Sum(elfCalories[len(elfCalories)-3:]...)
}

func (d *Day1) init() {
	for _, elf := range strings.Split(d.Input, "\n\n") {
		calorieCounts := []int{}
		for _, snack := range strings.Split(elf, "\n") {
			snackCalories, _ := strconv.Atoi(snack)
			calorieCounts = append(calorieCounts, snackCalories)
		}
		d.elves = append(d.elves, calorieCounts)
	}
}
