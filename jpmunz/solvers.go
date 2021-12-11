package jpmunz

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"advent2021/jpmunz/day01"
	"advent2021/jpmunz/day02"
	"advent2021/jpmunz/day03"
	"advent2021/jpmunz/day04"
	"advent2021/jpmunz/day05"
	"advent2021/jpmunz/day06"
	"advent2021/jpmunz/day07"
	"advent2021/jpmunz/day08"
	"advent2021/lib"
)

type Solvers struct{}

func (s *Solvers) GetSolver(day int) (lib.Solutions, error) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

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
