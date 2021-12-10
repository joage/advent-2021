package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/4
type Solution struct{}

var (
	space = regexp.MustCompile(`\s+`)
)

type cell struct {
	marked bool
	number int
}

type line []*cell

type board struct {
	rows    []line
	columns []line
	cells   []*cell
}

func (s Solution) Part1(lines []string) (string, error) {
	numbers, boards, err := parseInput(lines)
	if err != nil {
		return "", err
	}

	for _, n := range numbers {
		for _, b := range boards {
			b.mark(n)
			if b.hasWin() {
				log.Info().Int("winning_number", n).Str("winning_board", b.str()).
					Int("unmarked_total", b.sumUnmarked()).Msg("Found winner")
				return strconv.Itoa(b.sumUnmarked() * n), nil
			}
		}
	}

	log.Info().Msg("No winner found")
	return "", nil
}

func (s Solution) Part2(lines []string) (string, error) {
	numbers, boards, err := parseInput(lines)
	if err != nil {
		return "", err
	}

	losingBoards := make(map[int]board)
	for i, b := range boards {
		losingBoards[i] = b
	}

	for _, n := range numbers {
		for i, b := range boards {
			b.mark(n)

			if b.hasWin() {
				delete(losingBoards, i)

				if len(losingBoards) == 0 {
					log.Info().Int("winning_number", n).Int("winning_board_number", i+1).Str("winning_board", b.str()).
						Int("unmarked_total", b.sumUnmarked()).Msg("Found winner")

					return strconv.Itoa(b.sumUnmarked() * n), nil
				}
			}
		}
	}

	log.Info().Msg("No winner found")
	return "", nil
}

func (b *board) str() string {
	var rows []string

	/*
		for _, r := range b.rows {
			var row []string

			for _, c := range r {
				if c.marked {
					row = append(row, fmt.Sprintf("*%d*", c.number))
				} else {
					row = append(row, fmt.Sprintf(" %d ", c.number))
				}
			}
			rows = append(rows, strings.Join(row, " "))
		}

	*/

	for i, _ := range b.rows {
		var row []string

		for _, c := range b.columns {
			if c[i].marked {
				row = append(row, fmt.Sprintf("*%d*", c[i].number))
			} else {
				row = append(row, fmt.Sprintf(" %d ", c[i].number))
			}
		}
		rows = append(rows, strings.Join(row, " "))
	}

	return strings.Join(rows, "\n")
}

func (b *board) mark(n int) {
	for _, c := range b.cells {
		if c.number == n {
			c.marked = true
		}
	}
}

func (b *board) sumUnmarked() int {
	var sum int
	for _, c := range b.cells {
		if !c.marked {
			sum += c.number
		}
	}

	return sum
}

func (b *board) hasWin() bool {
	for _, r := range b.rows {
		if r.hasWin() {
			return true
		}
	}

	for _, c := range b.columns {
		if c.hasWin() {
			return true
		}
	}

	return false
}

func (l line) hasWin() bool {
	for _, c := range l {
		if !c.marked {
			return false
		}
	}

	return true
}

func buildBoard(input []string) (board, error) {
	var b board

	for i, l := range input {
		nl := space.ReplaceAllString(l, " ")

		tokens := strings.Split(nl, " ")

		if i == 0 {
			b.columns = make([]line, len(tokens))
			b.rows = make([]line, len(input))
		}

		for j, t := range tokens {
			n, err := strconv.Atoi(t)
			if err != nil {
				return board{}, err
			}
			c := &cell{
				marked: false,
				number: n,
			}

			b.rows[i] = append(b.rows[i], c)
			b.columns[j] = append(b.columns[j], c)
			b.cells = append(b.cells, c)
		}
	}

	return b, nil
}

func parseInput(lines []string) ([]int, []board, error) {
	var numbers []int
	var boards []board

	tokens := strings.Split(lines[0], ",")
	for _, t := range tokens {
		n, err := strconv.Atoi(t)
		if err != nil {
			return nil, nil, err
		}
		numbers = append(numbers, n)
	}

	for i := 2; i <= len(lines); i += 6 {
		b, err := buildBoard(lines[i : i+5])
		if err != nil {
			return nil, nil, err
		}
		boards = append(boards, b)
	}

	for _, b := range boards {
		log.Debug().Msg(b.str())
	}

	return numbers, boards, nil
}
