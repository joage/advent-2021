package day09

import (
	"testing"

	"advent2021/lib"
)

func TestDay9(t *testing.T) {
	lines, err := lib.ReadLines("../inputs/day9.txt")
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
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			expected: "15",
		},
		{
			name:     "actual test",
			input:    lines,
			expected: "508",
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
