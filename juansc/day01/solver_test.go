package day01

import (
	"testing"

	"advent2021/lib"
)

func TestDay1Part1(t *testing.T) {
	lines, err := lib.ReadLines("../inputs/day1.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name: "example",
			input: []string{
				"199",
				"200",
				"208",
				"210",
				"200",
				"207",
				"240",
				"269",
				"260",
				"263",
			},
			expected: "7",
		},
		{
			name:     "actual test",
			input:    lines,
			expected: "1665",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			solver := Day1{}
			actual, err := solver.Part1(test.input)
			if err != nil {
				t.Errorf("encountered error creating solution: %v", err)
			}
			if actual != test.expected {
				t.Errorf("did not match expected: actual=%s expected=%s", actual, test.expected)
			}
		})
	}
}

func TestDay1Part2(t *testing.T) {
	lines, err := lib.ReadLines("../inputs/day1.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name: "example",
			input: []string{
				"199",
				"200",
				"208",
				"210",
				"200",
				"207",
				"240",
				"269",
				"260",
				"263",
			},
			expected: "5",
		},
		{
			name:     "actual test",
			input:    lines,
			expected: "1702",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			solver := Day1{}
			actual, err := solver.Part2(test.input)
			if err != nil {
				t.Errorf("encountered error creating solution: %v", err)
			}
			if actual != test.expected {
				t.Errorf("did not match expected: actual=%s expected=%s", actual, test.expected)
			}
		})
	}
}
