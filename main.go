package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/nesselchen/aoc-2024/aoc"
	"github.com/nesselchen/aoc-2024/aoc/input"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Please pass at least one day to solve")
		os.Exit(1)
	}
	for _, arg := range flag.Args() {
		day, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Println("Day argument is ", day, " not a number")
		}

		fmt.Print("Day ", arg, ": ")

		// determine solver
		s, found := aoc.Solvers[int(day)]
		if !found {
			fmt.Println("-")
			continue
		}

		// load and split data
		path := fmt.Sprintf("aoc/day%d/input.txt", day)
		f, err := os.Open(path)
		if err != nil {
			fmt.Println("failed to read input file")
			continue
		}
		defer f.Close()

		sc := bufio.NewScanner(f)
		var lines input.Lines
		for sc.Scan() {
			lines = append(lines, []byte(sc.Text()))
		}

		// run solver
		fst := s.First(lines.Copy()) // copy to allow mutation on the input
		fmt.Print(fst)
		if s, ok := s.(aoc.PartTwoSolver); ok {
			snd := s.Second(lines)
			fmt.Print(" | ", snd)
		}
		fmt.Println()
	}
}
