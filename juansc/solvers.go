package juansc

import (
	"fmt"

	"advent2021/juansc/day01"
	"advent2021/juansc/day02"
	"advent2021/juansc/day03"
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
	default:
		return nil, fmt.Errorf("no solver exists for the given day")
	}
}
