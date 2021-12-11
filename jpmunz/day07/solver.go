package day07

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/7
type Solution struct{}

func (s Solution) Part1(lines []string) (string, error) {
	positions, err := parsePositions(lines)
	if err != nil {
		return "", err
	}

	// https://www.calculatorsoup.com/calculators/statistics/mean-median-mode.php
	var median int
	if len(positions)%2 == 0 {
		median = positions[(len(positions)-1)/2]
	} else {
		m1 := positions[(len(positions) / 2)]
		m2 := positions[(len(positions)/2)-1]
		median = (m1 + m2) / 2
	}

	var fuelSpent int
	for _, n := range positions {
		fuelSpent += int(math.Abs(float64(n) - float64(median)))
	}

	return strconv.Itoa(fuelSpent), nil
}

type alignment struct {
	position    int
	consumption int
}

func (s Solution) Part2(lines []string) (string, error) {
	positions, err := parsePositions(lines)
	if err != nil {
		return "", err
	}

	min := positions[0]
	max := positions[len(positions)-1]
	best := alignment{}

	for x := min; x <= max; x++ {
		var consumption int
		for _, p := range positions {
			consumption += calculateConsumption(p, x)
		}

		if best.consumption == 0 || consumption <= best.consumption {
			best.consumption = consumption
			best.position = x
		}
	}

	log.Debug().Int("position", best.position).Msg("Best found")

	return strconv.Itoa(best.consumption), nil
}

func calculateConsumption(x1 int, x2 int) int {
	n := math.Abs(float64(x1) - float64(x2))
	r := int(((n * (n - 1)) / 2) + n)

	log.Debug().Int("x1", x1).Int("x2", x2).Int("n", int(n)).Int("result", r).Msg("Calculated Consumption")

	return r
}

func parsePositions(lines []string) ([]int, error) {
	var positions []int
	for _, t := range strings.Split(lines[0], ",") {
		n, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}
		positions = append(positions, n)
	}

	sort.Ints(positions)

	return positions, nil
}
