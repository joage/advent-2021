package day04

import (
	"testing"

	"advent2021/lib"
)

func TestPart1(t *testing.T) {
	exampleLines, err := lib.ReadLines("../inputs/day4example.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	lines, err := lib.ReadLines("../inputs/day4.txt")
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
			input: exampleLines,
			expected: "4512",
		},
		{
			name:     "actual test",
			input:    lines,
			expected: "unknown",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			solver := Solution{}
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

func TestPart2(t *testing.T) {
	exampleLines, err := lib.ReadLines("../inputs/day4example.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	lines, err := lib.ReadLines("../inputs/day4.txt")
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
			input: exampleLines,
			expected: "1924",
		},
		{
			name:     "actual test",
			input:    lines,
			expected: "13158",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			solver := Solution{}
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
