package day04

import (
	"strconv"
	"strings"
)

func ConvertToRange(s string) []int {
	start, end := getStartEnd(s)
	var span []int
	for i := start; i <= end; i++ {
		span = append(span, i)
	}
	return span
}

func getStartEnd(s string) (int, int) {
	parts := strings.Split(s, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return start, end
}
