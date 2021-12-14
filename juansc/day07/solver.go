package day07

import (
	"math"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (string, error) {
	input := strings.Split(lines[0], ",")
	locations := make([]int, len(input))
	for i, l := range input {
		locations[i], _ = strconv.Atoi(l)
	}
	crabs := newCrabLocations(locations)
	minCost := crabs.findMinCost(constantCost)
	return strconv.Itoa(minCost), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	input := strings.Split(lines[0], ",")
	locations := make([]int, len(input))
	for i, l := range input {
		locations[i], _ = strconv.Atoi(l)
	}
	crabs := newCrabLocations(locations)
	minCost := crabs.findMinCost(triangularCost)
	return strconv.Itoa(minCost), nil
}

type crabLocations struct {
	// Count the number of crabs at a given location
	total          int
	min            int
	max            int
	locationCounts map[int]int
}

func newCrabLocations(locations []int) crabLocations {
	loc := map[int]int{}
	min := math.MaxInt
	max := math.MinInt
	for _, l := range locations {
		val, _ := loc[l]
		loc[l] = val + 1
		if l > max {
			max = l
		}
		if l < min {
			min = l
		}
	}
	return crabLocations{locationCounts: loc, total: len(locations), min: min, max: max}
}

func (c crabLocations) findMinCost(fn costFn) int {
	minCost := math.MaxInt
	for i := c.min; i <= c.max; i++ {
		cost := c.costToLocation(i, fn)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

type costFn func(dist int) int

func constantCost(dist int) int {
	return dist
}

func triangularCost(dist int) int {
	return (dist + 1) * dist / 2
}

func (c crabLocations) costToLocation(targetLocation int, fn costFn) int {
	cost := 0
	for loc, num := range c.locationCounts {
		dist := targetLocation - loc
		if dist < 0 {
			dist *= -1
		}
		// The cost to moving to this location is the product of the distance
		// and the number of crabs that need to go there
		cost += fn(dist) * num
	}
	return cost
}