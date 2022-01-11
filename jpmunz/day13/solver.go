package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/13
type Solution struct{}

type point struct {
	x int
	y int
}

func (p point) str() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

type fold struct {
	direction string
	line      int
}

func (f fold) str() string {
	return fmt.Sprintf("%v=%v", f.direction, f.line)
}

type paper struct {
	points map[point]bool
	width  int
	height int
}

func buildPaper(lines []string) (paper, []fold, error) {
	p := paper{
		points: make(map[point]bool),
	}
	var folds []fold
	var parsingFolds bool

	for _, l := range lines {
		if l == "" {
			parsingFolds = true
			continue
		}

		if parsingFolds {
			parts := strings.Split(l, " ")
			instruction := strings.Split(parts[2], "=")
			line, err := strconv.Atoi(instruction[1])
			if err != nil {
				return p, folds, err
			}

			folds = append(folds, fold{
				direction: instruction[0],
				line:      line,
			})
		} else {
			coord := strings.Split(l, ",")
			x, err := strconv.Atoi(coord[0])
			if err != nil {
				return p, folds, err
			}

			y, err := strconv.Atoi(coord[1])
			if err != nil {
				return p, folds, err
			}

			p.points[point{x: x, y: y}] = true

			if x > p.width {
				p.width = x
			}

			if y > p.height {
				p.height = y
			}
		}
	}

	return p, folds, nil
}

func (p *paper) str() string {
	var rows []string
	for y := 0; y <= p.height; y++ {
		var row []string
		for x := 0; x <= p.width; x++ {
			isDot, ok := p.points[point{x: x, y: y}]
			if ok && isDot {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
		}

		rows = append(rows, strings.Join(row, ""))
	}

	return strings.Join(rows, "\n")
}

func (p *paper) applyFold(f fold) {
	l := 1
	var remove []point

	if f.direction == "y" {
		for {
			foldFrom := f.line + l
			foldTo := f.line - l

			if foldFrom > p.height {
				break
			}

			for x := 0; x <= p.width; x++ {
				from := point{x: x, y: foldFrom}
				to := point{x: x, y: foldTo}

				p.points[to] = p.points[to] || p.points[from]
			}

			l++
		}

		p.height = f.line - 1

		for k, _ := range p.points {
			if k.y >= f.line {
				remove = append(remove, k)
			}
		}
	} else {
		for {
			foldFrom := f.line + l
			foldTo := f.line - l

			if foldFrom > p.width {
				break
			}

			for y := 0; y <= p.height; y++ {
				from := point{x: foldFrom, y: y}
				to := point{x: foldTo, y: y}

				p.points[to] = p.points[to] || p.points[from]
			}

			l++
		}

		p.width = f.line - 1

		for k, _ := range p.points {
			if k.x >= f.line {
				remove = append(remove, k)
			}
		}
	}

	for _, pt := range remove {
		delete(p.points, pt)
	}
}

func (s Solution) Part1(lines []string) (string, error) {
	p, folds, err := buildPaper(lines)
	if err != nil {
		return "", err
	}

	log.Debug().Msg("\n" + p.str())

	f := folds[0]
	log.Debug().Str("fold", f.str()).Msg("Processing fold")
	p.applyFold(f)
	log.Debug().Msg("\n" + p.str())

	var totalDots int
	for _, isDot := range p.points {
		if isDot {
			totalDots++
		}
	}

	return strconv.Itoa(totalDots), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	p, folds, err := buildPaper(lines)
	if err != nil {
		return "", err
	}

	log.Debug().Msg("\n" + p.str())

	for _, f := range folds {
		log.Debug().Str("fold", f.str()).Msg("Processing fold")
		p.applyFold(f)
		log.Debug().Msg("\n" + p.str())
	}

	return "\n" + p.str(), nil
}
