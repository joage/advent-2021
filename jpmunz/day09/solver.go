package day09

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/9
type Solution struct{}

type point struct {
	x int
	y int
}

func (p point) str() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

type heightmap struct {
	points map[point]int
}

func buildHeightMap(lines []string) (heightmap, error) {
	h := heightmap{
		points: make(map[point]int),
	}

	for y, l := range lines {
		xPoints := strings.Split(l, "")
		for x, s := range xPoints {
			n, err := strconv.Atoi(s)
			if err != nil {
				return h, err
			}

			h.points[point{x: x, y: y}] = n
		}
	}

	return h, nil
}

func getNeighbours(p point) []point {
	return []point{
		{
			x: p.x - 1,
			y: p.y,
		},
		{
			x: p.x + 1,
			y: p.y,
		},
		{
			x: p.x,
			y: p.y - 1,
		},
		{
			x: p.x,
			y: p.y + 1,
		},
	}
}

func (h heightmap) isLowPoint(p point) bool {
	for _, n := range getNeighbours(p) {
		height, ok := h.points[n]
		if ok && height <= h.points[p] {
			return false
		}
	}

	return true
}

func (h heightmap) calculateBasinSize(p point, seen map[point]bool) int {
	if seen == nil {
		seen = make(map[point]bool)
	}
	var inBasin []point
	for _, n := range getNeighbours(p) {
		height, ok := h.points[n]
		_, wasSeen := seen[n]
		if !wasSeen && ok && height < 9 && height > h.points[p] {
			inBasin = append(inBasin, n)
			seen[n] = true
		}
	}

	var size int

	log.Debug().Str("basin_point", p.str()).Int("height", h.points[p]).Msg("calculating basin")
	for _, b := range inBasin {
		size += h.calculateBasinSize(b, seen)
	}

	return 1 + size
}

func (s Solution) Part1(lines []string) (string, error) {
	h, err := buildHeightMap(lines)
	if err != nil {
		return "", err
	}

	var sum int
	for p, height := range h.points {
		if h.isLowPoint(p) {
			log.Debug().Str("point", p.str()).Msg("Found low point")
			sum += height + 1
		}
	}

	return strconv.Itoa(sum), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	h, err := buildHeightMap(lines)
	if err != nil {
		return "", err
	}

	var basinSizes []int
	for p, _ := range h.points {
		if h.isLowPoint(p) {
			log.Debug().Str("point", p.str()).Msg("Found low point")
			basinSizes = append(basinSizes, h.calculateBasinSize(p, nil))
		}
	}

	log.Debug().Ints("basin_sizes", basinSizes).Msg("Found basins")
	sort.Ints(basinSizes)

	total := basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]

	return strconv.Itoa(total), nil
}
