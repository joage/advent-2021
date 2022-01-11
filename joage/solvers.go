package joage

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"advent2021/joage/day01"
	"advent2021/joage/day02"
	"advent2021/joage/day03"
	"advent2021/lib"
)

type Solvers struct{}

func (s *Solvers) GetSolver(day int) (lib.Solutions, error) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	switch day {
	case 1:
		return &day01.Day1{}, nil
	case 2:
		return &day02.Day2{}, nil
	case 3:
		return &day03.Solution{}, nil
	default:
		return nil, fmt.Errorf("no solver exists for the given day")
	}
}
