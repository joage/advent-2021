package day01

import (
	"fmt"
	"strconv"
)

type Day1 struct {}

func (d Day1) Part1(lines []string) (string, error) {
	var err error
	input := make([]int, len(lines))
	for i, line := range lines {
		input[i], err = strconv.Atoi(line)
		if err != nil {
			return "", fmt.Errorf("error parsing line %d as number (%s): %v", i, line, err)
		}
	}
	count := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			count++
		}
	}
	return strconv.Itoa(count), nil

	/*
	// Part 2
	smooth_count := 0
	for i := 0; i < len(input)-3; i++ {
		if input[i] < input[i+3] { // subtract input[i+1] + input[i+2] from both sides
			smooth_count++
		}
	}
	fmt.Println("2) smooth count: ", smooth_count)

	 */
}

func (d Day1) Part2(lines []string) (string, error) {
	panic("not implemented")
}
