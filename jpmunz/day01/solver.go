package day01

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/1
type Solution struct{}

func (s Solution) Part1(lines []string) (string, error) {
	measurements, err := getMeasurements(lines)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(getIncreases(measurements)), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	measurements, err := getMeasurements(lines)
	if err != nil {
		return "", err
	}

	var sliding []int
	for i, _ := range measurements {
		if i+2 < len(measurements) {
			sliding = append(sliding, measurements[i]+measurements[i+1]+measurements[i+2])
		}
	}

	return strconv.Itoa(getIncreases(sliding)), nil
}

func getMeasurements(lines []string) ([]int, error) {
	var measurements []int
	for _, t := range lines {
		i, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}

		measurements = append(measurements, i)
	}

	return measurements, nil
}

func getIncreases(l []int) int {
	prevM := 0
	increased := 0
	for _, m := range l {
		if prevM > 0 && m > prevM {
			log.Debug().Msg(fmt.Sprintf("%d (increased)", m))
			increased++
		} else {
			log.Debug().Msg(fmt.Sprintf("%d", m))
		}

		prevM = m
	}

	return increased
}
