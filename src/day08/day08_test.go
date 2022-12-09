package day08_test

import (
	_ "embed"
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/src/day08"
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
		expectedPartOne: 21,
	},
	{
		input:           input,
		expectedPartOne: 1698,
	},
}

func TestNewGrid(t *testing.T) {
	g := day08.NewGrid(testInput)
	assert.Equal(t, g.Get(0, 0), 3)
	assert.Equal(t, g.Get(1, 1), 5)
	assert.Equal(t, g.Get(2, 2), 3)
	assert.Equal(t, g.Get(4, 3), 9)
}

func TestHidden(t *testing.T) {
	g := day08.NewGrid(testInput)
	// Outer edge is always visible
	assert.Equal(t, day08.IsVisible(g, 0, 0), true)
	assert.Equal(t, day08.IsVisible(g, 3, 4), true)
	assert.Equal(t, day08.IsVisible(g, 4, 4), true)

	assert.Equal(t, day08.IsVisible(g, 1, 1), true)
	assert.Equal(t, day08.IsVisible(g, 2, 1), true)
	assert.Equal(t, day08.IsVisible(g, 3, 1), false)

	assert.Equal(t, day08.IsVisible(g, 1, 2), true)
	assert.Equal(t, day08.IsVisible(g, 2, 2), false)
	assert.Equal(t, day08.IsVisible(g, 3, 2), true)

	assert.Equal(t, day08.IsVisible(g, 2, 3), true)
	assert.Equal(t, day08.IsVisible(g, 3, 3), false)
	assert.Equal(t, day08.IsVisible(g, 1, 3), false)
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		g := day08.NewGrid(tc.input)
		v := 0
		for i, row := range g.Sparse() {
			for j := range row {
				if day08.IsVisible(g, j, i) {
					v += 1
				}
			}
		}
		assert.Equal(t, v, tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {}
