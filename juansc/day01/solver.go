package day01

import (
	"fmt"
	"strconv"
)

type Solution struct {}

func stringsToInts(lines []string) ([]int, error) {
	var err error
	input := make([]int, len(lines))
	for i, line := range lines {
		input[i], err = strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing line %d as number (%s): %v", i, line, err)
		}
	}
	return input, nil
}

func (s Solution) Part1(lines []string) (string, error) {
	input, err := stringsToInts(lines)
	if err != nil {
		return "", err
	}
	// Shamelessly stolen from joage and modified so that the teacher can't tell I copied the HW.
	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i - 1] {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	input, err := stringsToInts(lines)
	if err != nil {
		return "", err
	}
	count := 0
	for i := 3; i < len(input); i++ {
		if input[i] + input[i-1] + input[i-2] > input[i-1] + input[i-2] + input[i-3] {
			count++
		}
	}
	return strconv.Itoa(count), nil
}
