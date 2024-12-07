package aoc

import (
	"github.com/nesselchen/aoc-2024/aoc/day4"
	"github.com/nesselchen/aoc-2024/aoc/input"
)

type BaseSolver interface {
	SolveFirst(input.Lines) int
}

type PartTwoSolver interface {
	BaseSolver
	SolveSecond(input.Lines) int
}

var Solvers = map[int]BaseSolver{
	4: day4.Solver{},
}
