package day09

import (
	"fmt"
	"sort"
	"strconv"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (string, error) {
	floorHeights := newHeights(lines)
	lowPoints := floorHeights.lowPoints()
	riskTotal := 0
	for _, p := range lowPoints {
		fmt.Println(p.height)
		riskTotal += p.height + 1
	}
	fmt.Println(len(lowPoints))
	return strconv.Itoa(riskTotal), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	floorHeights := newHeights(lines)
	basinID := 0
	for i, p := range floorHeights.grid {
		// Points of max height cannot be part of a basin
		if p.height == 9 {
			continue
		}
		// The point may have already been assigned to a basin
		if p.basinID != -1 {
			continue
		}
		floorHeights.assignBasin(i, basinID)
		basinID++
		// Can we pass some recursive function?
	}
	basinSizes := make([]int, basinID)
	for _, p := range floorHeights.grid {
		if p.height == 9 {
			continue
		}
		basinSizes[p.basinID]++
	}
	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})
	return strconv.Itoa(basinSizes[0] * basinSizes[1] * basinSizes[2]), nil
}

type floorPoint struct {
	ind        int
	height     int
	skip       bool
	isLowPoint bool
	basinID    int
}

type heights struct {
	grid   []floorPoint
	height int
	width  int
}

func (h *heights) assignBasin(ind, basinID int) {
	point := h.grid[ind]
	if point.height == 9 {
		return
	}
	if point.basinID != -1 {
		return
	}
	h.grid[ind].basinID = basinID
	for _, neighbor := range h.getNeighbors(ind) {
		index := neighbor.ind
		h.assignBasin(index, basinID)
	}
}

func newHeights(lines []string) heights {
	h := heights{
		grid:   []floorPoint{},
		height: len(lines),
		width:  0,
	}
	ind := 0
	for _, line := range lines {
		h.width = len(line)
		for _, heightStr := range line {
			height, _ := strconv.Atoi(string(heightStr))
			h.grid = append(h.grid, floorPoint{
				height: height,
				ind:    ind,
				// Use -1 as the unassigned basin ID
				basinID: -1,
			})
			ind++
		}
	}
	return h
}

// getNeighbors returns a slice of pointers for the neighbors of the points
// at ind. Note that ind is the 1-d index in the grid
func (h *heights) getNeighbors(ind int) []*floorPoint {
	points := []*floorPoint{}
	// The point is not in the top row
	if ind/h.width != 0 {
		points = append(points, &h.grid[ind-h.width])
	}
	// The point is not on the leftmost col
	if ind%h.width != 0 {
		points = append(points, &h.grid[ind-1])
	}
	// The point is not on the rightmost col
	if ind%h.width != h.width-1 {
		points = append(points, &h.grid[ind+1])
	}
	// The point is not on the bottom row
	if ind/h.width != h.height-1 {
		points = append(points, &h.grid[ind+h.width])
	}
	return points
}

func (h *heights) lowPoints() []*floorPoint {
	arrLowPoints := []*floorPoint{}
	for i, currentPoint := range h.grid {
		// It is possible that analyzing one of this point's neighbors
		// disqualified the current point as a low point.
		if currentPoint.skip {
			continue
		}
		neighbors := h.getNeighbors(i)
		isLow := true
		for _, neighbor := range neighbors {
			// Since the neighbor is a low point, by definition this current point
			// cannot be a low point. Mark it as skipped!
			if neighbor.isLowPoint {
				isLow = false
				currentPoint.skip = true
				break
			}
			if currentPoint.height > neighbor.height {
				isLow = false
				currentPoint.skip = true
				break
			}
			if neighbor.height == currentPoint.height {
				// If the current point has the same height
				// as its neighbor, neither is a low point.
				// Mark them both as needing to be skipped
				isLow = false
				currentPoint.skip = true
				neighbor.skip = true
			}
		}
		currentPoint.isLowPoint = isLow
		h.grid[i] = currentPoint
		if isLow {
			arrLowPoints = append(arrLowPoints, &h.grid[i])
		}
	}
	return arrLowPoints
}
