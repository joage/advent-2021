package day08

import (
"testing"

"advent2021/lib"
	. "github.com/onsi/gomega"
)

func TestPart1(t *testing.T) {
	exampleLines, err := lib.ReadLines("../inputs/day8_example.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	lines, err := lib.ReadLines("../inputs/day8.txt")
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
			expected: "38",
		},
		{
			name:     "actual test",
			input:    lines,
			expected: "344835",
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
	exampleLines, err := lib.ReadLines("../inputs/day8_example.txt")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}
	lines, err := lib.ReadLines("../inputs/day8.txt")
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
			expected: "61229",
		},
		{
			name:     "actual test",
			input:    lines,
			expected: "1027483",
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

func TestNewDecoder(t *testing.T) {
	g := NewGomegaWithT(t)
	decoder, err := newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(decoder.encode(segmentA)).To(Equal(segmentD))
	g.Expect(decoder.encode(segmentB)).To(Equal(segmentE))
	g.Expect(decoder.encode(segmentC)).To(Equal(segmentA))
	g.Expect(decoder.encode(segmentD)).To(Equal(segmentF))
	g.Expect(decoder.encode(segmentE)).To(Equal(segmentG))
	g.Expect(decoder.encode(segmentF)).To(Equal(segmentB))
	g.Expect(decoder.encode(segmentG)).To(Equal(segmentC))

	g.Expect(decoder.identifyNumber("acedgfb")).To(Equal(8))
	g.Expect(decoder.identifyNumber("cdfbe")).To(Equal(5))
	g.Expect(decoder.identifyNumber("gcdfa")).To(Equal(2))
	g.Expect(decoder.identifyNumber("fbcad")).To(Equal(3))
	g.Expect(decoder.identifyNumber("dab")).To(Equal(7))
	g.Expect(decoder.identifyNumber("cefabd")).To(Equal(9))
	g.Expect(decoder.identifyNumber("cdfgeb")).To(Equal(6))
	g.Expect(decoder.identifyNumber("eafb")).To(Equal(4))
	g.Expect(decoder.identifyNumber("cagedb")).To(Equal(0))
	g.Expect(decoder.identifyNumber("ab")).To(Equal(1))
}

