package day08

import (
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/array"
	"github.com/mleone10/advent-of-code-2022/pkg/grid"
)

type Day08 struct {
	Input string
}

func NewGrid(input string) grid.Plane[int] {
	g := grid.Plane[int]{}
	for i, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for j, col := range strings.Split(row, "") {
			height, _ := strconv.Atoi(col)
			g.Set(j, i, height)
		}
	}
	return g
}

func IsVisible(g grid.Plane[int], x, y int) bool {
	if x == 0 || y == 0 || x == g.Width()-1 || y == g.Height()-1 {
		// Point is the edge of the map, thus is visible
		return true
	}
	return isVisibleFromRight(g, x, y) ||
		isVisibleFromLeft(g, x, y) ||
		isVisibleFromTop(g, x, y) ||
		isVisibleFromBottom(g, x, y)
}

func isVisibleFromTop(g grid.Plane[int], x, y int) bool {
	return array.Max(g.Col(x)[:y]) < g.Get(x, y)
}

func isVisibleFromBottom(g grid.Plane[int], x, y int) bool {
	return array.Max(g.Col(x)[y+1:]) < g.Get(x, y)
}

func isVisibleFromLeft(g grid.Plane[int], x, y int) bool {
	return array.Max(g.Row(y)[:x]) < g.Get(x, y)
}

func isVisibleFromRight(g grid.Plane[int], x, y int) bool {
	return array.Max(g.Row(y)[x+1:]) < g.Get(x, y)
}
