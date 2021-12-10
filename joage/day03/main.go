package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan()
}

func scan() {
	file, err := os.Open("report.in")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// part 1
	var digits [12]int

	// count ones and zeros, a one means add 1 to the digit, 0 means subtract 1
	s := bufio.NewScanner(file)
	for s.Scan() {
		if err != nil {
			fmt.Println(err)
		}
		line := strings.Split(s.Text(), "")
		for i, digit := range line {
			one, err := strconv.ParseBool(digit)
			if err != nil {
				fmt.Println(err)
			}
			if one {
				digits[i] += 1
			} else {
				digits[i] -= 1
			}
		}
	}
	//fmt.Println(digits)

	// generate binary numbers
	mostCommon, leastCommon := "", ""
	for _, i := range digits {
		if i > 0 {
			mostCommon = mostCommon + "1"
			leastCommon = leastCommon + "0"
		} else if i < 0 {
			mostCommon = mostCommon + "0"
			leastCommon = leastCommon + "1"
		} else {
			panic(fmt.Sprintf("there is no most common value for the bit at index %d", i))
		}
	}

	gamma, err := strconv.ParseInt(mostCommon, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	epsilon, err := strconv.ParseInt(leastCommon, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("part 1...")
	fmt.Printf("gamma: binary '%s' decimal %d\n", mostCommon, gamma)
	fmt.Printf("epsilon: binary '%s' decimal %d\n", leastCommon, epsilon)
	fmt.Printf("power consumption is %d\n", gamma*epsilon)

}
