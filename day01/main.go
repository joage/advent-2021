package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)


func main() {
	input, err := scan()
	if err != nil {
		fmt.Println("error processing input: ", err)
	  	return
	}
	// Part 1
	count := 0
	for i := 0; i < len(input) - 1; i++ {
		if input[i] < input[i+1] {
			count++
		}
	}
	fmt.Println("1) scan count: ", count)

	// Part 2
	smooth_count := 0
	for i := 0; i < len(input) - 3; i++ {
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
		lines= append(lines, read)
	}

	return lines, s.Err()
}
