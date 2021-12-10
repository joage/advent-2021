package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Solver interface {
	Solve(lines []string) (string, error)
}

type Day1Part1 struct {}

func (d Day1Part1) Solve(lines []string) (string, error) {
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

func main() {
	input, err := scan()
	if err != nil {
		fmt.Println("error processing input: ", err)
		return
	}
	// Part 1
	count := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			count++
		}
	}
	fmt.Println("1) scan count: ", count)

	// Part 2
	smooth_count := 0
	for i := 0; i < len(input)-3; i++ {
		if input[i] < input[i+3] { // subtract input[i+1] + input[i+2] from both sides
			smooth_count++
		}
	}
	fmt.Println("2) smooth count: ", smooth_count)
}

// read file input
func scan() ([]int, error) {
	file, err := os.Open("depths.in")
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	var lines []int
	s := bufio.NewScanner(file)
	for s.Scan() {
		read, err := strconv.Atoi(s.Text())
		if err != nil {
			return []int{}, err
		}
		lines = append(lines, read)
	}

	return lines, s.Err()
}
