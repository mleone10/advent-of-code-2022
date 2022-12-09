package day09_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/pkg/grid"
	"github.com/mleone10/advent-of-code-2022/src/day09"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	expectedPartOne int
	expectedPartTwo int
}{
	{
		input:           testInput,
		expectedPartOne: 13,
	},
	{
		input:           input,
		expectedPartOne: 6236,
	},
}

func TestMoveHead(t *testing.T) {
	r := day09.Rope{}
	r.MoveHead(day09.Dir["D"], 3)
	assert.Equal(t, r.Head.X, 0)
	assert.Equal(t, r.Head.Y, 3)
	r.MoveHead(day09.Dir["U"], 3)
	assert.Equal(t, r.Head.X, 0)
	assert.Equal(t, r.Head.Y, 0)
	r.MoveHead(day09.Dir["L"], 3)
	assert.Equal(t, r.Head.X, -3)
	assert.Equal(t, r.Head.Y, 0)
	r.MoveHead(day09.Dir["R"], 3)
	assert.Equal(t, r.Head.X, 0)
	assert.Equal(t, r.Head.Y, 0)
}

func TestMoveHeadMovesTail(t *testing.T) {
	r := day09.Rope{}
	r.MoveHead(day09.Dir["D"], 3)
	assert.Equal(t, r.Head.X, 0)
	assert.Equal(t, r.Head.Y, 3)
	assert.Equal(t, r.Tail.X, 0)
	assert.Equal(t, r.Tail.Y, 2)

	r.MoveHead(day09.Dir["U"], 3)
	assert.Equal(t, r.Head.X, 0)
	assert.Equal(t, r.Head.Y, 0)
	assert.Equal(t, r.Tail.X, 0)
	assert.Equal(t, r.Tail.Y, 1)

	r.MoveHead(day09.Dir["L"], 3)
	assert.Equal(t, r.Head.X, -3)
	assert.Equal(t, r.Head.Y, 0)
	assert.Equal(t, r.Tail.X, -2)
	assert.Equal(t, r.Tail.Y, 0)

	r.MoveHead(day09.Dir["R"], 3)
	assert.Equal(t, r.Head.X, 0)
	assert.Equal(t, r.Head.Y, 0)
	assert.Equal(t, r.Tail.X, -1)
	assert.Equal(t, r.Tail.Y, 0)
}

func TestMoveHeadDiagonal(t *testing.T) {
	r := day09.Rope{Head: grid.Point{X: 1, Y: -1}}
	r.MoveHead(day09.Dir["U"], 1)
	assert.Equal(t, r.Head.X, 1)
	assert.Equal(t, r.Head.Y, -2)
	assert.Equal(t, r.Tail.X, 1)
	assert.Equal(t, r.Tail.Y, -1)
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		r := day09.Rope{}
		r.SimulateMoves(strings.Split(strings.TrimSpace(tc.input), "\n"))
		assert.Equal(t, len(r.TailPositions()), tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {}
