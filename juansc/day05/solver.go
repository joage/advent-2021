package day05

import (
	"math"
	"strconv"
	"strings"
)

type Solution struct{}

type point struct {
	x int
	y int
}

func newPoint(input string) point {
	parts := strings.Split(input, ",")
	p := point{}
	p.x, _ = strconv.Atoi(parts[0])
	p.y, _ = strconv.Atoi(parts[1])
	return p

}

func (p point) eq(other point) bool {
	return p.x == other.x && p.y == other.y
}

type latticeSlope int

const (
	slopeHorizontal  = 0
	slopeVertical    = 1
	slopePosDiagonal = 2
	slopeNegDiagonal = 3
)

// a segmnt is normalized so that a.x <= b.x. If there is a tie then a.y < b.y.
type segment struct {
	a     point
	b     point
	slope latticeSlope
}

func newSegment(input string) segment {
	parts := strings.Split(input, " -> ")
	a := newPoint(parts[0])
	b := newPoint(parts[1])
	if a.eq(b) {
		panic("points are the same, segment invalid")
	}
	var slope latticeSlope
	if a.x == b.x {
		// Slope is vertical since the x matches
		slope = slopeVertical
		// normalize points
		if a.y > b.y {
			a, b = b, a
		}
		//
	} else if a.y == b.y {
		// Slope is vertical since the y matches
		slope = slopeHorizontal
		// normalize points
		if a.x > b.x {
			a, b = b, a
		}
	} else if math.Abs(float64(a.y - b.y)) == math.Abs(float64(a.x - b.x)) {
		// since |dx| = |dy|, then the slope is diagonal
		// normalize points. At this point we know they don't have x's or y's. When we normalize
		// we make sure that the one with the smallest x comes first.
		if a.x > b.x {
			a, b = b, a
		}
		slope = slopePosDiagonal
		if a.y > b.y {
			slope = slopeNegDiagonal
		}
	} else {
		// We have an unsupported slope
		panic("segment is not on the lattice")
	}
	return segment{a: a, b:b, slope: slope}
}

// latticePoints returns all the points from segment.a to segment.b that fall on the
// integer lattice. This implemntation assumes that lines are either horizontal, vertical,
// or have a slope of 1 or -1.
func (s segment) latticePoints() []point {
	var dx, dy int
	switch s.slope {
	case slopeHorizontal:
		dx, dy = 1, 0
	case slopeVertical:
		dx, dy = 0, 1
	case slopePosDiagonal:
		dx, dy = 1, 1
	case slopeNegDiagonal:
		dx, dy = 1, -1
	default:
		panic("slope is not on a lattice")
	}
	points := []point{}
	currentPoint := s.a
	for !currentPoint.eq(s.b) {
		points = append(points, currentPoint)
		currentPoint.x += dx
		currentPoint.y += dy
	}
	points = append(points, s.b)
	return points
}

func (s Solution) Part1(lines []string) (string, error) {
	return calculateOverlappingPoints(lines, false)
}

func (s Solution) Part2(lines []string) (string, error) {
	return calculateOverlappingPoints(lines, true)
}

func calculateOverlappingPoints(lines []string, withDiagnoals bool) (string, error) {
	lattice := map[point]int{}
	for _, line := range lines {
		segment := newSegment(line)
		if !withDiagnoals && (segment.slope == slopeNegDiagonal || segment.slope == slopePosDiagonal) {
			continue
		}
		for _, point := range segment.latticePoints() {
			count, _ := lattice[point]
			lattice[point] = count + 1
		}
	}
	// count the number of points that were visitied more than once
	total := 0
	for _, count := range lattice {
		if count > 1 {
			total += 1
		}
	}
	return strconv.Itoa(total), nil
}
