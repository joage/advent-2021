package day01

import (
	"fmt"
	"strconv"
)

type Day1 struct{}

func (d Day1) Part1(lines []string) (string, error) {
	depths, err := convertToInts(lines)
	if err != nil {
		return "", fmt.Errorf("error converting input to integers: %v", err)
	}
	count := 0
	for i := 0; i < len(depths)-1; i++ {
		if depths[i] < depths[i+1] {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func (d Day1) Part2(lines []string) (string, error) {
	depths, err := convertToInts(lines)
	if err != nil {
		return "", fmt.Errorf("error converting input to integers: %v", err)
	}
	smoothCount := 0
	for i := 0; i < len(depths)-3; i++ {
		if depths[i] < depths[i+3] { // subtract input[i+1] + input[i+2] from both sides
			smoothCount++
		}
	}
	return strconv.Itoa(smoothCount), nil
}

func convertToInts(lines []string) ([]int, error) {
	var err error
	input := make([]int, len(lines))
	for i, line := range lines {
		input[i], err = strconv.Atoi(line)
		if err != nil {
			return []int{}, fmt.Errorf("error parsing line %d as number (%s): %v", i, line, err)
		}
	}
	return input, nil
}
