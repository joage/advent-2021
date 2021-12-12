package day06

import (
	"testing"

	"advent2021/lib"
	. "github.com/onsi/gomega"
)

func TestPart1(t *testing.T) {
	exampleLines, err := lib.ReadLines("../inputs/day6_example.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	lines, err := lib.ReadLines("../inputs/day6.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "example",
			input:    exampleLines,
			expected: "5934",
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
	exampleLines, err := lib.ReadLines("../inputs/day6_example.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	lines, err := lib.ReadLines("../inputs/day6.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "example",
			input:    exampleLines,
			expected: "26984457539",
		},
		{
			name:     "actual test",
			input:    lines,
			expected: "1632146183902",
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

func TestNewCohorts(t *testing.T) {
	tests := []struct {
		name     string
		inputs   []int
		expected cohorts
	}{
		{
			name:   "empty",
			inputs: []int{},
			expected: cohorts{
				numTicks:     0,
				cohortCounts: [numDays]int{},
			},
		},
		{
			name:   "all same day",
			inputs: []int{3, 3, 3},
			expected: cohorts{
				numTicks: 0,
				cohortCounts: [numDays]int{
					0, 0, 0, 3,
					0, 0, 0, 0, 0,
				},
			},
		},
		{
			name:   "first three days",
			inputs: []int{1, 2, 3, 2, 1},
			expected: cohorts{
				numTicks: 0,
				cohortCounts: [numDays]int{
					0, 2, 2, 1,
					0, 0, 0, 0, 0,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := newCohorts(test.inputs)
			if actual.cohortCounts != test.expected.cohortCounts {
				t.Errorf("cohort counts do not match %v != %v", actual.cohortCounts, test.expected.cohortCounts)
			}
		})
	}
}

func TestUpdateCohorts(t *testing.T) {
	// "3,4,3,1,2"
	tests := []struct {
		name     string
		initial   cohorts
		expectedForTicks []cohorts
	}{
		{
			name:   "ticks",
			initial: cohorts{
				numTicks:     0,
				cohortCounts: [9]int{
					0,
					1,
					1,
					2,
					1,
					0,
					0,
					0,
					0,
				},
			},
			expectedForTicks: []cohorts{
				{
					numTicks: 1,
					cohortCounts: [9]int{
						1,
						1,
						2,
						1,
						0,
						0,
						0,
						0,
						0,
					},
				},
				{
					numTicks: 2,
					cohortCounts: [9]int{
						1,
						2,
						1,
						0,
						0,
						0,
						1,
						0,
						1,
					},
				},
				{
					numTicks: 3,
					cohortCounts: [9]int{
						2,
						1,
						0,
						0,
						0,
						1,
						1,
						1,
						1,
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			for _, expected := range test.expectedForTicks {
				test.initial.update()
				g.Expect(test.initial).To(Equal(expected))
			}
		})
	}
}
