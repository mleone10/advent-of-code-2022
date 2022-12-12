package day09

import (
	"log"
	"strconv"
	"strings"

	"github.com/mleone10/advent-of-code-2022/pkg/grid"
	"github.com/mleone10/advent-of-code-2022/pkg/linkedlist"
	"github.com/mleone10/advent-of-code-2022/pkg/maputil"
)

type Dir grid.Point

var (
	DirUp    = Dir{X: 0, Y: -1}
	DirDown  = Dir{X: 0, Y: 1}
	DirLeft  = Dir{X: -1, Y: 0}
	DirRight = Dir{X: 1, Y: 0}
)

var Cmds = map[string]Dir{
	"U": DirUp,
	"D": DirDown,
	"L": DirLeft,
	"R": DirRight,
}

type Knot struct {
	Pos     grid.Point
	visited map[grid.Point]bool
}

func NewRope(len int) *linkedlist.Node[*Knot] {
	head := linkedlist.NewNode(newKnot())
	for i := 0; i < len-1; i++ {
		head.Tail().LinkNext(linkedlist.NewNode(newKnot()))
	}
	return head
}

func newKnot() *Knot {
	return &Knot{visited: map[grid.Point]bool{{X: 0, Y: 0}: true}}
}

func SimulateMoves(r *linkedlist.Node[*Knot], mvs []string) {
	for _, mv := range mvs {
		mParts := strings.Split(mv, " ")
		dist, _ := strconv.Atoi(mParts[1])
		MoveN(r, Cmds[mParts[0]], dist)
	}
}

func MoveN(r *linkedlist.Node[*Knot], d Dir, dist int) {
	for i := 0; i < dist; i++ {
		moveHead(r, d)
	}
}

func moveHead(r *linkedlist.Node[*Knot], d Dir) {
	r.Value().Pos.X += d.X
	r.Value().Pos.Y += d.Y
	r.Value().Visit()
	log.Println(r.Value().Pos)

	if r.Next() != nil {
		updateKnot(r.Next())
	}
}

func updateKnot(k *linkedlist.Node[*Knot]) {
	var movedX, movedY bool
	dx := k.Prev().Value().Pos.X - k.Value().Pos.X
	dy := k.Prev().Value().Pos.Y - k.Value().Pos.Y

	if abs(dx) >= 2 {
		k.Value().Pos.X += (dx / abs(dx))
		if abs(dy) != 0 {
			k.Value().Pos.Y += (dy / abs(dy))
		}
		movedX = true
	}

	if abs(dy) >= 2 {
		k.Value().Pos.Y += (dy / abs(dy))
		if abs(dx) != 0 {
			k.Value().Pos.X += (dx / abs(dx))
		}
		movedY = true
	}

	// If we moved at all, remember the new position.
	if movedX || movedY {
		k.Value().Visit()
	}

	// If there's another knot, update it's position as well.
	if k.Next() != nil {
		updateKnot(k.Next())
	}
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func (k *Knot) Visit() {
	k.visited[grid.Point{X: k.Pos.X, Y: k.Pos.Y}] = true
}

func (k Knot) Visited() []grid.Point {
	return maputil.Keys(k.visited)
}
