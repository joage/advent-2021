package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/11
type Solution struct{}

type point struct {
	x int
	y int
}

type octopus struct {
	energy     int
	hasFlashed bool
}

func (p point) str() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

type cavern struct {
	points map[point]*octopus
	width  int
	height int
}

func buildCavern(lines []string) (cavern, error) {
	c := cavern{
		points: make(map[point]*octopus),
		height: len(lines),
	}

	for y, l := range lines {
		xPoints := strings.Split(l, "")
		c.width = len(xPoints)
		for x, s := range xPoints {
			n, err := strconv.Atoi(s)
			if err != nil {
				return c, err
			}

			c.points[point{x: x, y: y}] = &octopus{energy: n}
		}
	}

	return c, nil
}

func getNeighbours(p point) []point {
	return []point{
		// topleft
		{
			x: p.x - 1,
			y: p.y - 1,
		},
		// top
		{
			x: p.x,
			y: p.y - 1,
		},
		// topright
		{
			x: p.x + 1,
			y: p.y - 1,
		},
		// left
		{
			x: p.x - 1,
			y: p.y,
		},
		// right
		{
			x: p.x + 1,
			y: p.y,
		},
		// bottomleft
		{
			x: p.x - 1,
			y: p.y + 1,
		},
		// bottom
		{
			x: p.x,
			y: p.y + 1,
		},
		// bottomright
		{
			x: p.x + 1,
			y: p.y + 1,
		},
	}
}

func (c *cavern) incr(p point) int {
	var flashes int
	o, ok := c.points[p]
	if !ok {
		return 0
	}
	o.energy += 1

	if o.energy > 9 && !o.hasFlashed {
		o.hasFlashed = true
		flashes += 1
		for _, n := range getNeighbours(p) {
			flashes += c.incr(n)
		}
	}

	return flashes
}

func (c cavern) str() string {
	grid := make([]string, c.height)

	for y := 0; y < c.height; y++ {
		line := make([]string, c.width)
		for x := 0; x < c.width; x++ {
			line[x] = strconv.Itoa(c.points[point{x: x, y: y}].energy)
		}
		grid[y] = strings.Join(line, "")
	}

	return strings.Join(grid, "\n")
}

func (c *cavern) simulateStep() int {
	var flashes int
	for p := range c.points {
		flashes += c.incr(p)
	}

	for _, o := range c.points {
		if o.energy > 9 {
			o.energy = 0
			o.hasFlashed = false
		}
	}

	return flashes
}

func (s Solution) Part1(lines []string) (string, error) {
	c, err := buildCavern(lines)
	if err != nil {
		return "", err
	}

	var totalFlashes int
	log.Debug().Msg("\n" + c.str())
	for s := 1; s <= 100; s++ {
		flashesForStep := c.simulateStep()
		log.Debug().Int("step", s).Int("flashes", flashesForStep).Msg("Completed step")
		log.Debug().Msg("\n" + c.str())
		totalFlashes += flashesForStep
	}

	return strconv.Itoa(totalFlashes), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	c, err := buildCavern(lines)
	if err != nil {
		return "", err
	}

	var totalFlashes int
	var step int
	log.Debug().Msg("\n" + c.str())
	for {
		step++
		flashesForStep := c.simulateStep()
		if flashesForStep == len(c.points) {
			break
		}

		log.Debug().Int("step", step).Msg("Completed step")
		log.Debug().Msg("\n" + c.str())
		totalFlashes += flashesForStep
	}

	return strconv.Itoa(step), nil
}
