package day03

import (
	"strings"
)

type Day03 struct {
	Input string
}

func GroupCompartments(contents string) []string {
	return []string{contents[:len(contents)/2], contents[len(contents)/2:]}
}

func FindCommonContents(compartments []string) string {
	common := compartments[0]

	var newCommon []rune
	for _, comp := range compartments[1:] {
		for _, char := range comp {
			if Contains([]rune(common), char) {
				newCommon = append(newCommon, char)
			}
		}
		common = string(newCommon)
		newCommon = nil
	}

	return common
}

// TODO: move this to array package
func Contains[T comparable](arr []T, find T) bool {
	for _, v := range arr {
		if v == find {
			return true
		}
	}
	return false
}

func CalculatePriority(b byte) int {
	if string(b) == strings.ToLower(string(b)) {
		return int(b) - 96
	}
	return int(b) - 38
}
