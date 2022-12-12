package day12_test

import (
	_ "embed"
	"testing"
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
		expectedPartOne: 31,
	},
	{
		input:           input,
		expectedPartOne: 0,
	},
}

func TestSolvePartOne(t *testing.T) {}

func TestSolvePartTwo(t *testing.T) {}
