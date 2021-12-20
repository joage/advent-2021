package day09

type Solution struct {}

func (s Solution) Part1(lines []string) (string, error) {
	panic("implement me")
}

func (s Solution) Part2(lines []string) (string, error) {
	panic("implement me")
}

type floorPoint struct {
	height int
	skip bool
	isLowPoint bool
}

type heights struct {
	grid []floorPoint
	height int
	width int
}

// getNeighbors returns a slice of pointers for the neighbors of the points
// at ind. Note that ind is the 1-d index in the grid
func (h *heights) getNeighbors(ind int) []*floorPoint {
	points := []*floorPoint{}
	// The point is not in the top row
	if ind / h.width != 0 {
		points = append(points, &h.grid[ind - h.width])
	}
	// The point is not on the leftmost col
	if ind % h.width != 0 {
		points = append(points, &h.grid[ind - 1])
	}
	// The point is not on the rightmost col
	if ind % h.width != h.width - 1 {
		points = append(points, &h.grid[ind + 1])
	}
	// The point is not on the bottom row
	if ind / h.width != h.height - 1 {
		points = append(points, &h.grid[ind + h.width])
	}
	return points
}

func (h *heights) lowPoints() []int {
	for i, p := range h.grid {
		// It is possible that analyzing one of this current point's neighbors
		// disqualified it as a low point.
		if p.skip {
			continue
		}
		neighbors := h.getNeighbors(i)
		for _, neighbor := range neighbors {
			// Since the neighbor is a low point, by definition this current point
			// cannot be a low point. Mark it as skipped!
			if neighbor.isLowPoint {
				p.skip = true
				h.grid[i] = p
				break
			}
			if neighbor.height == p.height {
				// If the current point has the same height
				// as its neighbor, neither is a low point.
				// Mark them both as needing to be skipped
				p.skip = true
				neighbor.skip = true
			}
		}

	}
}