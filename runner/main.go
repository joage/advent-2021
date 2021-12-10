package main

import (
	"fmt"
	"time"

	"advent2021/joage/day01"
	"advent2021/lib"
)

type Solver interface {
	Solve([]string) (string, error)
}

func main() {
	day := ""
	var solver1, solver2 Solver
	lines, err := lib.ReadLines(fmt.Sprintf("input/day%s", day))
	if err !=nil {
		panic("could not read file")
	}
	switch day {
	case "1":
		solver1 = day01.Day1Part1{}
		solver2 = day01.Day1Part1{}
	default:
		fmt.Println("solution coming soon!")
		return
	}
	start := time.Now()
	solution1, err := solver1.Solve(lines)
	if err != nil {
		panic("???")
	}
	fmt.Println("solution for day 1", solution1)
	solution2, err := solver2.Solve(lines)
	if err != nil {
		panic("???")
	}
	fmt.Println("solution for day 1", solution2)
	fmt.Println("Ran all solutions in %v", time.Since(start))
}
