package day09

import (
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/grid"
)

var Dir = map[string]grid.Point{
	"U": {X: 0, Y: -1},
	"D": {X: 0, Y: 1},
	"L": {X: -1, Y: 0},
	"R": {X: 1, Y: 0},
}

type Rope struct {
	Head    grid.Point
	Tail    grid.Point
	visited grid.Plane[bool]
}

func (r *Rope) SimulateMoves(moves []string) {
	r.visited.Set(0, 0, true)
	for _, m := range moves {
		mParts := strings.Split(m, " ")
		dist, _ := strconv.Atoi(mParts[1])
		r.MoveHead(Dir[mParts[0]], dist)
	}
}

func (r *Rope) MoveHead(dir grid.Point, dist int) {
	for i := 0; i < dist; i++ {
		r.step(dir)
	}
}

func (r *Rope) step(dir grid.Point) {
	prevHead := r.Head
	r.Head.X += dir.X
	r.Head.Y += dir.Y
	if !r.Head.Neighboring(r.Tail) {
		r.Tail = prevHead
		r.visited.Set(r.Tail.X, r.Tail.Y, true)
	}
}

func (r Rope) TailPositions() []grid.Point {
	var points []grid.Point
	for i, row := range r.visited.Sparse() {
		for j := range row {
			points = append(points, grid.Point{X: j, Y: i})
		}
	}
	return points
}
