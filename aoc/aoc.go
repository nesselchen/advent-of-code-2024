package aoc

import (
	"github.com/nesselchen/aoc-2024/aoc/day2"
	"github.com/nesselchen/aoc-2024/aoc/day3"
	"github.com/nesselchen/aoc-2024/aoc/day4"
	"github.com/nesselchen/aoc-2024/aoc/day5"
	"github.com/nesselchen/aoc-2024/aoc/input"
)

type BaseSolver interface {
	First(input.Lines) int
}

type PartTwoSolver interface {
	BaseSolver
	Second(input.Lines) int
}

var Solvers = map[int]BaseSolver{
	2: day2.Solver{},
	3: day3.Solver{},
	4: day4.Solver{},
	5: day5.Solver{},
}
