package day1

import (
	"slices"

	"github.com/nesselchen/aoc-2024/aoc/input"
)

type Solver struct{}

func (Solver) First(lines input.Lines) int {
	var totalDistances int
	left, right := parse(lines)
	slices.Sort(left)
	slices.Sort(right)
	for i := range left {
		totalDistances += absDiff(left[i], right[i])
	}
	return totalDistances
}

func parse(lines input.Lines) ([]int, []int) {
	var left, right []int
	for _, line := range lines {
		ln, rn, _ := input.SplitOnce(string(line), "   ")
		left = append(left, input.ParseInt(ln))
		right = append(right, input.ParseInt(rn))
	}
	return left, right
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
