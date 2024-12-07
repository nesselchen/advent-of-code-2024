package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/nesselchen/aoc-2024/aoc"
	"github.com/nesselchen/aoc-2024/aoc/input"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please pass the day you want to solve")
		os.Exit(1)
	}
	arg := os.Args[1]
	day, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		fmt.Println("Day argument is not a number")
		os.Exit(1)
	}

	// determine solver
	s, found := aoc.Solvers[int(day)]
	if !found {
		fmt.Println("Seems like haven't started solving day", day, "yet")
	}

	// load and split data
	path := fmt.Sprintf("aoc/day%d/data.txt", day)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to read input file")
		os.Exit(1)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines input.Lines
	for sc.Scan() {
		lines = append(lines, []byte(sc.Text()))
	}

	// run solver
	fmt.Printf("Day %d: ", day)
	fst := s.SolveFirst(lines)
	fmt.Print(fst)

	if s, ok := s.(aoc.PartTwoSolver); ok {
		snd := s.SolveSecond(lines)
		fmt.Print(" / ", snd)
	}
	//
	fmt.Println()
}
