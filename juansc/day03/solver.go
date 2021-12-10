package day03

import (
	"fmt"
	"strconv"
)

type Solution struct{}

type binaryTracker struct {
	digits []int
	processedInputs int
}

func newBinaryTracker(input string) binaryTracker {
	b := binaryTracker{digits: make([]int, len(input))}
	b.update(input)
	return b
}

func (b *binaryTracker) update(input string) {
	for i, c := range input {
		if c == '1' {
			b.digits[i]++
		}
	}
	b.processedInputs++
}

func (b *binaryTracker) gamma() int {
	numLiteral := ""
	for _, v := range b.digits {
		// 1 was the most common bit
		if v > b.processedInputs / 2 {
			numLiteral += "1"
		} else {
			numLiteral += "0"
		}
	}
	return parseBinaryString(numLiteral)
}

func (b *binaryTracker) epsilon() int {
	numLiteral := ""
	for _, v := range b.digits {
		// 1 was the most common bit
		if v > b.processedInputs / 2 {
			numLiteral += "0"
		} else {
			numLiteral += "1"
		}
	}
	return parseBinaryString(numLiteral)
}

func parseBinaryString(bin string) int {
	val, _ := strconv.ParseInt(bin, 2, 64)
	return int(val)

}

func (s Solution) Part1(lines []string) (string, error) {
	if len(lines) == 0 {
		return "", fmt.Errorf("lines may not be empty")
	}
	tracker := newBinaryTracker(lines[0])
	for i := 1; i < len(lines); i ++ {
		tracker.update(lines[i])
	}
	return strconv.Itoa(tracker.gamma() * tracker.epsilon()), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	tree := newbinaryTree(lines)
	currentNode := tree.head
	stringValue := ""
	for currentNode.HasChildren() {
		var newStr string
		newStr, currentNode = currentNode.MostCommonChild()
		stringValue += newStr
		if currentNode == nil {
			break
		}
	}
	oxygenRating := parseBinaryString(stringValue)

	currentNode = tree.head
	stringValue = ""
	for currentNode.HasChildren() {
		var newStr string
		newStr, currentNode = currentNode.LeastCommonChild()
		stringValue += newStr
		if currentNode == nil {
			break
		}
	}
	co2Rating := parseBinaryString(stringValue)

	return strconv.Itoa(oxygenRating * co2Rating), nil
}
