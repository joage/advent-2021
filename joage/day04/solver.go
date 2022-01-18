package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (string, error) {
	var digits [12]int
	for _, line := range lines {
		parts := strings.Split(line, "")
		for i, digit := range parts {
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
	mostCommon, leastCommon := "", ""
	for _, i := range digits {
		if i > 0 {
			mostCommon = mostCommon + "1"
			leastCommon = leastCommon + "0"
		} else if i < 0 {
			mostCommon = mostCommon + "0"
			leastCommon = leastCommon + "1"
		} else {
			return "", fmt.Errorf("there is no most common value for the bit at index %d", i)
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
	log.Info().Int64("decimal", gamma).Str("binary", mostCommon).Msg("gamme")
	log.Info().Int64("decimal", epsilon).Str("binary", leastCommon).Msg("epsilon")
	return strconv.Itoa(int(gamma * epsilon)), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	panic("not implememnted")
}
