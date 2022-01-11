package day02

import (
	"strconv"
	"strings"
)

type Day2 struct{}

func (d Day2) Part1(lines []string) (string, error) {
	depth, horizontal := 0, 0

	for _, line := range lines {
		elements := strings.Split(line, " ")
		amount, err := strconv.Atoi(elements[1])
		if err != nil {
			return "", err
		}
		switch elements[0] {
		case "forward":
			horizontal += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		default:
			panic("bad command")
		}
	}
	return strconv.Itoa(depth * horizontal), nil
}

func (d Day2) Part2(lines []string) (string, error) {
	// part 2 vars
	depth, horizontal, aim := 0, 0, 0

	for _, line := range lines {
		elements := strings.Split(line, " ")
		amount, err := strconv.Atoi(elements[1])
		if err != nil {
			return "", err
		}
		switch elements[0] {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		default:
			panic("bad command")
		}
	}

	return strconv.Itoa(depth * horizontal), nil

}
