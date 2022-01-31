package day09

import (
	"sort"
	"strconv"
)

type Solution struct{}

type tube struct {
	depths    []int
	height    int
	width     int
	lowPoints map[int]bool
}

func newLavaTube(lines []string) tube {
	t := tube{
		width: len(lines[0]),
		height: len(lines),
		depths: make([]int, len(lines)*len(lines[0])),
		lowPoints: map[int]bool{},
	}
	i := 0
	for _, line := range lines {
		for _, unit := range line {
			depth, _ := strconv.Atoi(string(unit))
			t.depths[i] = depth
			i++
		}
	}
	return t
}

func (t tube) getNeighbors(index int) []int {
	var neighbors []int
	col := index % t.width
	row := index / t.width
	if row > 0 {
		neighbors = append(neighbors, index-t.width)
	}
	if row < t.height-1 { // actually the height
		neighbors = append(neighbors, index+t.width)
	}
	if col > 0 {
		neighbors = append(neighbors, index-1)
	}
	if col < t.width-1 {
		neighbors = append(neighbors, index+1)
	}
	return neighbors
}

func (t tube) isLow(index int) bool {
	// exit early if we have already determined this to not be a low point
	if isLow, seen := t.lowPoints[index]; seen && !isLow {
		return false
	}

	neighbors := t.getNeighbors(index)
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
	neighbors := t.getNeighbors(lowPoint)

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
