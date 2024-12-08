package day5

import (
	"slices"
	"strings"

	"github.com/nesselchen/aoc-2024/aoc/input"
)

type Solver struct{}

// Part 1

func (Solver) First(lines input.Lines) int {
	rules, pages := parseInput(lines)
	return solveAll(rules, pages)
}

func solveAll(rules map[int][]int, pages [][]int) int {
	var sum int
	for _, p := range pages {
		if valid(rules, p) {
			sum += int(p[len(p)/2])
		}
	}
	return sum
}

func valid(rules map[int][]int, page []int) bool {
	for i, n := range page {
		for _, bad := range rules[n] {
			if slices.Contains(page[:i], bad) {
				return false
			}
		}
	}
	return true
}

// Part 2

func (Solver) Second(lines input.Lines) int {
	rules, pages := parseInput(lines)
	return checkAndFixAll(rules, pages)
}

func checkAndFixAll(rules map[int][]int, pages [][]int) int {
	var sum int
	for _, p := range pages {
		if valid(rules, p) {
			continue
		}
		fix(rules, p)
		sum += int(p[len(p)/2])
	}
	return sum
}

var first = false

func fix(rules map[int][]int, page []int) []int {
	for i := range page {
		n := page[0]
		for j := 0; j < len(page)-i-1; j++ {
			if slices.Contains(rules[n], page[j+1]) {
				n = page[j+1]
				continue
			}
			page[j], page[j+1] = page[j+1], page[j]
		}
	}
	return page
}

// Used by both

func parseInput(lines input.Lines) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	pagesStart := 0
	for i, l := range lines {
		if len(l) == 0 {
			pagesStart = i + 1
			break
		}
		fst, snd, _ := input.SplitOnce(string(l), "|")
		before := input.ParseInt(fst)
		after := input.ParseInt(snd)
		rules[before] = append(rules[before], after)
	}
	pages := make([][]int, 0)
	for i, l := range lines[pagesStart:] {
		splits := strings.Split(string(l), ",")
		pages = append(pages, make([]int, 0))
		for _, s := range splits {
			n := input.ParseInt(s)
			pages[i] = append(pages[i], n)
		}
	}
	return rules, pages
}
