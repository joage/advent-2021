package juansc

import (
	"fmt"

	"advent2021/juansc/day01"
	"advent2021/juansc/day02"
	"advent2021/juansc/day03"
	"advent2021/juansc/day04"
	"advent2021/juansc/day05"
	"advent2021/juansc/day06"
	"advent2021/juansc/day07"
	"advent2021/juansc/day08"
	"advent2021/lib"
)

type Solvers struct{}

func (s *Solvers) GetSolver(day int) (lib.Solutions, error) {
	switch day {
	case 1:
		return &day01.Solution{}, nil
	case 2:
		return &day02.Solution{}, nil
	case 3:
		return &day03.Solution{}, nil
	case 4:
		return &day04.Solution{}, nil
	case 5:
		return &day05.Solution{}, nil
	case 6:
		return &day06.Solution{}, nil
	case 7:
		return &day07.Solution{}, nil
	case 8:
		return &day08.Solution{}, nil
	default:
		return nil, fmt.Errorf("no solver exists for the given day")
	}
}
