package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/2
type Solution struct{}

type position struct {
	horizontal int
	depth      int
	aim        int
}

func (s Solution) Part1(lines []string) (string, error) {
	var p position
	for _, t := range lines {
		tokens := strings.Split(t, " ")
		i, err := strconv.Atoi(tokens[1])
		if err != nil {
			return "", err
		}

		switch tokens[0] {
		case "forward":
			p.horizontal += i
		case "down":
			p.depth += i
		case "up":
			p.depth -= i
		}
	}

	log.Info().Msg(fmt.Sprintf("Final position: %v", p))

	return strconv.Itoa(p.horizontal * p.depth), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	var p position
	for _, t := range lines {
		tokens := strings.Split(t, " ")
		i, err := strconv.Atoi(tokens[1])
		if err != nil {
			return "", err
		}

		switch tokens[0] {
		case "forward":
			p.horizontal += i
			p.depth += p.aim * i
		case "down":
			p.aim += i
		case "up":
			p.aim -= i
		}
	}

	log.Info().Msg(fmt.Sprintf("Final position: %v", p))

	return strconv.Itoa(p.horizontal * p.depth), nil
}
