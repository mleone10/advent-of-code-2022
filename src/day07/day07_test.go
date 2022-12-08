package day07_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-2022/pkg/array"
	"github.com/mleone10/advent-of-code-2022/pkg/assert"
	"github.com/mleone10/advent-of-code-2022/src/day07"
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
		expectedPartOne: 95437,
	},
}

func TestChangeDir(t *testing.T) {
	d := day07.Day07{}
	d.ChangeDir("/")
	assert.Equal(t, d.Pwd(), "/")
	d.ChangeDir("a")
	assert.Equal(t, d.Pwd(), "/a")
	d.ChangeDir("..")
	assert.Equal(t, d.Pwd(), "/")
}

func TestDiscover(t *testing.T) {
	d := day07.Day07{}
	d.ChangeDir("/")
	d.Discover("dir a")
	d.Discover("123 b.txt")
	d.Discover("456 c.txt")
	assert.Contains(t, d.List(), "b.txt")
	assert.Contains(t, d.List(), "c.txt")
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		d := array.Reduce(strings.Split(strings.TrimSpace(tc.input), "\n"), func(d day07.Day07, cmd string) day07.Day07 {
			switch cmd[:4] {
			case "$ cd":
				args := strings.Split(cmd, " ")
				d.ChangeDir(args[2])
			case "$ ls":
			default:
				d.Discover(cmd)
			}
			return d
		}, day07.Day07{})
		assert.Equal(t, d.SumOfDirsMaxNBytes(100000), tc.expectedPartOne)
	}
}

func TestSolvePartTwo(t *testing.T) {}
