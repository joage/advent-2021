package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/5
type Solution struct{}

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func (s Solution) Part1(lines []string) (string, error) {
	return computeOverlap(lines, true)
}

func (s Solution) Part2(lines []string) (string, error) {
	return computeOverlap(lines, false)
}

func computeOverlap(input []string, ignoreDiagonals bool) (string, error) {
	var lines []line
	for _, l := range input {
		points := strings.Split(l, " -> ")
		lines = append(lines, line{start: pointFromString(points[0]), end: pointFromString(points[1])})
	}

	covered := make(map[point]int)
	for _, l := range lines {
		if ignoreDiagonals && l.diagonal() {
			continue
		}

		var points []string
		for _, p := range l.points() {
			points = append(points, p.str())
			covered[p]++
		}

		log.Debug().Strs("points", points).Msg(l.str())
	}

	var highOverlap int
	for y := 0; y < 10; y++ {
		var gridline []string
		for x := 0; x < 10; x++ {
			n := covered[point{x: x, y: y}]

			if n == 0 {
				gridline = append(gridline, ".")
			} else {
				gridline = append(gridline, strconv.Itoa(n))
			}
		}

		log.Debug().Msg(strings.Join(gridline, ""))
	}

	for _, v := range covered {
		if v >= 2 {
			highOverlap++
		}
	}

	return strconv.Itoa(highOverlap), nil
}

func minInt(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func (p point) str() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func (l line) str() string {
	return fmt.Sprintf("%v -> %v", l.start.str(), l.end.str())
}

func (l line) points() []point {
	var points []point

	x := l.start.x
	y := l.start.y

	/*
		x1 := minInt(l.start.x, l.end.x)
		x2 := maxInt(l.start.x, l.end.x)
		y1 := minInt(l.start.y, l.end.y)
		y2 := maxInt(l.start.y, l.end.y)
	*/

	for {
		points = append(points, point{x: x, y: y})

		if x == l.end.x && y == l.end.y {
			break
		}

		if x < l.end.x {
			x++
		} else if x > l.end.x {
			x--
		}

		if y < l.end.y {
			y++
		} else if y > l.end.y {
			y--
		}
	}

	return points
}

func (l line) diagonal() bool {
	return l.start.x != l.end.x && l.start.y != l.end.y
}

func pointFromString(s string) point {
	tokens := strings.Split(s, ",")

	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to convert string to point")
	}

	y, err := strconv.Atoi(tokens[1])
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to convert string to point")
	}

	return point{x: x, y: y}
}
