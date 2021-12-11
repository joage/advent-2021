package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/6
type Solution struct{}

type lanternFish struct {
	timer int
}

func (s Solution) Part1(lines []string) (string, error) {
	var fish []*lanternFish

	for _, t := range strings.Split(lines[0], ",") {
		n, err := strconv.Atoi(t)
		if err != nil {
			return "", err
		}
		fish = append(fish, &lanternFish{timer: n})
	}

	for i := 0; i < 80; i++ {
		var newFish []*lanternFish

		for _, f := range fish {
			nf := f.incrDay()

			if nf != nil {
				newFish = append(newFish, nf)
			}
		}

		fish = append(fish, newFish...)

		var population []int
		for _, f := range fish {
			population = append(population, f.timer)
		}

		log.Debug().Int("day", i+1).Ints("population", population).Msg("Day finished")
	}

	return strconv.Itoa(len(fish)), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	fishCounts := make(map[int]int)

	for _, t := range strings.Split(lines[0], ",") {
		n, err := strconv.Atoi(t)
		if err != nil {
			return "", err
		}
		fishCounts[n]++
	}

	var population int
	for d := 0; d < 256; d++ {
		newFishCounts := make(map[int]int)

		for i := 1; i <= 8; i++ {
			newFishCounts[i-1] = fishCounts[i]
		}

		newFishCounts[8] = fishCounts[0]
		newFishCounts[6] += fishCounts[0]

		fishCounts = newFishCounts
		population = 0
		var fishSummary []string
		for i := 0; i <= 8; i++ {
			population += fishCounts[i]
			fishSummary = append(fishSummary, fmt.Sprintf("%d:%d", i, fishCounts[i]))
		}

		log.Debug().Int("day", d+1).Int("population", population).Strs("summary", fishSummary).Msg("Day finished")
	}

	return strconv.Itoa(population), nil
}

func (l *lanternFish) incrDay() *lanternFish {
	if l.timer == 0 {
		l.timer = 6
		return &lanternFish{timer: 8}
	} else {
		l.timer--
		return nil
	}
}
