package day09

import (
	"sort"
	"strconv"
)

type Solution struct{}

type tube struct {
	depths    []int
	width     int
	lowPoints map[int]bool
}

func getNeighbors(width, index int) []int {
	var neighbors []int
	col := index % width
	row := index / width
	if row > 0 {
		neighbors = append(neighbors, index-width)
	}
	if row < width-1 { // actually the height
		neighbors = append(neighbors, index+width)
	}
	if col > 0 {
		neighbors = append(neighbors, index-1)
	}
	if col < width-1 {
		neighbors = append(neighbors, index+1)
	}
	return neighbors
}

func newLavaTube(lines []string) tube {
	t := tube{}
	t.width = len(lines[0])
	var depths []int
	for _, line := range lines {
		for _, unit := range line {
			depth, _ := strconv.Atoi(string(unit))
			depths = append(depths, depth)
		}
	}
	t.depths = depths
	t.lowPoints = map[int]bool{}
	return t
}

func (t tube) isLow(index int) bool {
	neighbors := getNeighbors(t.width, index)
	for _, n := range neighbors {
		if isNeighborLow, seen := t.lowPoints[n]; (seen && isNeighborLow) || (t.depths[n] <= t.depths[index]) {
			t.lowPoints[index] = false
			return false
		}
	}

	t.lowPoints[index] = true
	// index is a low point, all neighbors must not be low points
	for _, n := range neighbors {
		t.lowPoints[n] = false
	}
	return true
}

func (t tube) getBasinSizeIncluding(lowPoint int, visited map[int]bool) int {
	neighbors := getNeighbors(t.width, lowPoint)

	var upstream []int
	for _, n := range neighbors {
		if _, seen := visited[n]; !seen && t.depths[n] != 9 && t.depths[n] > t.depths[lowPoint] {
			upstream = append(upstream, n)
			visited[n] = true
		}
	}
	basinSize := 1
	for _, u := range upstream {
		basinSize += t.getBasinSizeIncluding(u, visited)
	}

	return basinSize
}

func (s Solution) Part1(lines []string) (string, error) {
	t := newLavaTube(lines)

	risk := 0
	for i, depth := range t.depths {
		if t.isLow(i) {
			risk += 1 + depth
		}
	}

	return strconv.Itoa(risk), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	t := newLavaTube(lines)

	var lowIndices []int
	for i := range t.depths {
		if t.isLow(i) {
			lowIndices = append(lowIndices, i)
		}
	}

	visited := map[int]bool{}
	var basinSizes []int
	for _, i := range lowIndices {
		basinSizes = append(basinSizes, t.getBasinSizeIncluding(i, visited))
	}
	sort.Ints(basinSizes)
	sliceLen := len(basinSizes)
	ans := basinSizes[sliceLen-1] * basinSizes[sliceLen-2] * basinSizes[sliceLen-3]
	return strconv.Itoa(ans), nil
}
