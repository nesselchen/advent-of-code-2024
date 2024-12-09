package day2

import (
	"slices"
	"strings"

	"github.com/nesselchen/aoc-2024/aoc/input"
)

type Solver struct{}

func (Solver) First(lines input.Lines) int {
	reports := parse(lines)
	var safe int
	for _, r := range reports {
		if len(r) <= 2 {
			safe++
			continue
		}
		ord := asc
		if r[0] > r[1] {
			ord = desc
		}
		if slices.IsSortedFunc(r, ord) {
			safe++
		}
	}
	return safe
}

func (Solver) Second(lines input.Lines) int {
	reports := parse(lines)
	var safe int
	for _, r := range reports {
		if len(r) < 2 {
			safe++
			continue
		}
		ord := asc
		if r[0] > r[1] {
			ord = desc
		}
		if slices.IsSortedFunc(r, ord) {
			safe++
			continue
		}
		for i := range r {
			removed := RemoveOne(r, i)
			ord := asc
			if removed[0] > removed[1] {
				ord = desc
			}
			if slices.IsSortedFunc(removed, ord) {
				safe++
				break
			}
		}
	}
	return safe
}

func parse(lines input.Lines) [][]int {
	reports := make([][]int, 0, lines.Height())
	for _, l := range lines {
		splits := strings.Split(string(l), " ")
		r := make([]int, 0, len(splits))
		for _, s := range splits {
			r = append(r, input.ParseInt(s))
		}
		reports = append(reports, r)
	}
	return reports
}

func isSafe(a, b int) bool {
	if diff := b - a; diff < 1 || diff > 3 {
		return false
	}
	return true
}

func asc(a, b int) int {
	if isSafe(b, a) {
		return 1
	}
	return -1
}

func desc(a, b int) int {
	if isSafe(a, b) {
		return 1
	}
	return -1
}

func RemoveOne(s []int, idx int) []int {
	removed := make([]int, 0, len(s)-1)
	for i, v := range s {
		if i == idx {
			continue
		}
		removed = append(removed, v)
	}
	return removed
}
