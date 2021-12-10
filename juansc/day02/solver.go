package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

type submarine struct {
	horizontal int
	depth      int
	aim        int
}

func (s Solution) Part1(lines []string) (string, error) {
	sub := submarine{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction, magnitudeStr := parts[0], parts[1]
		magnitude, err := strconv.Atoi(magnitudeStr)
		if err != nil {
			return "", fmt.Errorf("encountered error parsing %v", err)
		}
		switch strings.ToLower(direction) {
		case "forward":
			sub.horizontal += magnitude
		case "down":
			sub.depth += magnitude
		case "up":
			sub.depth -= magnitude
		}
	}
	return strconv.Itoa(sub.horizontal * sub.depth), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	sub := submarine{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction, magnitudeStr := parts[0], parts[1]
		magnitude, err := strconv.Atoi(magnitudeStr)
		if err != nil {
			return "", fmt.Errorf("encountered error parsing %v", err)
		}
		switch strings.ToLower(direction) {
		case "forward":
			sub.horizontal += magnitude
			sub.depth += sub.aim * magnitude
		case "down":
			sub.aim += magnitude
		case "up":
			sub.aim -= magnitude
		}
	}
	return strconv.Itoa(sub.horizontal * sub.depth), nil
}
