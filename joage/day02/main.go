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
	file, err := os.Open("commands.in")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// part 1 vars
	depth1, horizontal1 := 0, 0

	// part 2 vars
	depth2, horizontal2, aim := 0, 0, 0

	s := bufio.NewScanner(file)
	for s.Scan() {
		if err != nil {
			fmt.Println(err)
		}
		line := strings.Split(s.Text(), " ")
		dir := line[0]
		amount, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(err)
		}

		switch dir {
		case "forward":
			horizontal1 += amount
			horizontal2 += amount
			depth2 += aim * amount
		case "down":
			depth1 += amount
			aim += amount
		case "up":
			depth1 -= amount
			aim -= amount
		default:
			panic("bad command")
		}
	}

	fmt.Println("part 1: ", depth1*horizontal1)

	fmt.Println("part 2: ", depth2*horizontal2)

}
