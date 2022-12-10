package day09

import (
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/grid"
)

type Dir grid.Point

var (
	DirUp        = Dir{X: 0, Y: -1}
	DirDown      = Dir{X: 0, Y: 1}
	DirLeft      = Dir{X: -1, Y: 0}
	DirRight     = Dir{X: 1, Y: 0}
	DirUpRight   = Dir{X: 1, Y: -1}
	DirUpLeft    = Dir{X: -1, Y: -1}
	DirDownRight = Dir{X: 1, Y: 1}
	DirDownLeft  = Dir{X: -1, Y: 1}
)

var Cmds = map[string]Dir{
	"U": DirUp,
	"D": DirDown,
	"L": DirLeft,
	"R": DirRight,
}

type Rope struct {
	Head Knot
	Tail Knot
}

type Knot struct {
	Pos     grid.Point
	visited []grid.Point
	next    *Knot
	prev    *Knot
}

func NewRope(len int) Rope {
	head := newKnot()
	prev := &head
	tail := &head
	for i := 0; i < len-1; i++ {
		next := newKnot()
		prev.next = &next
		next.prev = prev
		tail = &next
		prev = &next
	}
	return Rope{Head: head, Tail: *tail}
}

func newKnot() Knot {
	return Knot{visited: []grid.Point{{X: 0, Y: 0}}}
}

func (r Rope) Length() int {
	return r.Head.length()
}

func (k Knot) length() int {
	if k.next == nil {
		return 1
	}
	return 1 + k.next.length()
}

func (r *Rope) SimulateMoves(moves []string) {
	for _, m := range moves {
		mParts := strings.Split(m, " ")
		dist, _ := strconv.Atoi(mParts[1])
		r.MoveHead(Cmds[mParts[0]], dist)
	}
}

func (r *Rope) MoveHead(dir Dir, dist int) {
	for i := 0; i < dist; i++ {
		r.Head.step(grid.Point(dir))
	}
}

func (k *Knot) step(next grid.Point) {
	prevPos := k.Pos
	// Move the current knot
	k.Pos.X = next.X
	k.Pos.Y = next.Y

	// Determine where to move the next knot
	if k.next != nil {
		kDir := k.chooseNext(prevPos)
		k.next.step(kDir)
	}
}

func (k Knot) chooseNext(prevPos grid.Point) grid.Point {
	return prevPos
}

// func (r *Rope) step(dir grid.Point) {
// 	prevHead := r.Head
// 	r.Head.X += dir.X
// 	r.Head.Y += dir.Y
// 	if !r.Head.Neighboring(r.Tail) {
// 		r.Tail = prevHead
// 		r.visited.Set(r.Tail.X, r.Tail.Y, true)
// 	}
// }

func (r Rope) TailPositions() []grid.Point {
	return r.Tail.visited
}
