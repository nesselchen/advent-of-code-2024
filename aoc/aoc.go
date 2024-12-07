package aoc

import "github.com/nesselchen/aoc-2024/aoc/day4"

type BaseSolver interface {
	SolveFirst([]string) int
}

type AdvancedSolver interface {
	BaseSolver
	SolveSecond([]string) int
}

var Solvers = map[int]BaseSolver{
	4: day4.Solver{},
	5: day4.Solver{},
}
