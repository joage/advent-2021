package day04

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	boardDim = 5
)

// This solution assumes that all boards eventually win if we read out all the numbers provided
// for the puzzle. If no board wins or there is a single board that does not win then the solution
// fails.


type Solution struct{}

type point struct {
	row int
	col int
}

type winStats struct {
	stepsToWin int
	score      int
}

func (s winStats) IsZero() bool {
	return s.stepsToWin == 0 && s.score == 0
}

type board struct {
	grid           []int
	numberLocation map[int]point
}

// evaluateWin consumes the numbers in the order the are called and returns a tuple with
// the (stepsToWin, score), where stepsToWin indicates how many numbers were called when
// the board won and score is calculated as n * sm, where n is the winning number and
// sm is the sum of all numbers that have not been called yet. A negative score indicates
// that the board never won.
// This method assumes that numbers are called only once.
func (b *board) evaluateWin(nums []int) (winStats, error) {
	numsMissingForRow := make([]int, boardDim)
	for i := range numsMissingForRow {
		numsMissingForRow[i] = boardDim
	}
	numsMissingForCol := make([]int, boardDim)
	for i := range numsMissingForCol {
		numsMissingForCol[i] = boardDim
	}
	// Calculate the sum of missing numbers. Each time we find one we will subtract a value.
	sumMissingNumbers := 0
	for _, n := range b.grid {
		sumMissingNumbers += n
	}

	for i, n := range nums {
		location, ok := b.numberLocation[n]
		if !ok {
			continue
		}
		// We found this number, so remove from the total
		sumMissingNumbers -= n
		// We found the last number for the row or column. Return score
		if numsMissingForCol[location.col] == 1 || numsMissingForRow[location.row] == 1 {
			return winStats{
				stepsToWin: i,
				score:      sumMissingNumbers * n,
			}, nil
		}
		numsMissingForCol[location.col] -= 1
		numsMissingForRow[location.row] -= 1
	}
	// The board never won
	return winStats{}, fmt.Errorf("board did not win")
}

func newBoard(lines []string) board {
	b := board{
		grid:           make([]int, 0, boardDim*boardDim),
		numberLocation: map[int]point{},
	}
	for row, line := range lines {
		col := 0
		for _, str := range strings.Split(line, " ") {
			if str != "" {
				n, _ := strconv.Atoi(str)
				b.grid = append(b.grid, n)
				b.numberLocation[n] = point{row: row, col: col}
				col += 1
			}
		}
	}
	return b
}

type game struct {
	// numbers that we will called out
	numbers []int
	boards  []board
}


func (g *game) findBoardScoreForStrategy(strategy boardStrategy) (int, error){
	boardScores := []winStats{}

	for _, board := range g.boards {
		// Only include boards that win. Keep in mind that this strategy assumes that
		// all boards eventually win. If there is one board that does not win or if there
		// are several boards that do not win there will be bugs.
		if stats, err := board.evaluateWin(g.numbers); err == nil {
			boardScores = append(boardScores, stats)
		}
	}
	if len(boardScores) == 0 {
		return -1, fmt.Errorf("no board won")
	}
	// Given all the board scores, pick the "best" one according to the given strategy
	strat := strategy(boardScores)
	sort.Slice(boardScores, strat)
	return boardScores[0].score, nil
}

// This returns a sorting function that we can use to find the board that was the "best" at something
// from a given arr.
type boardStrategy func(arr []winStats) func(i,j int) bool

var (
	winningBoardStrategy = func(arr []winStats) func(i, j int) bool {
		return func(i, j int) bool {return arr[i].stepsToWin < arr[j].stepsToWin}
	}
	losingBoardStrategy = func(arr []winStats) func(i, j int) bool {
		return func(i, j int) bool {return arr[i].stepsToWin > arr[j].stepsToWin}
	}
)

func newGame(lines []string) game {
	g := game{numbers: []int{}, boards: []board{}}
	for _, s := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(s)
		g.numbers = append(g.numbers, n)
	}

	// After the input line each board is preceeded by a new line and 5 lines.
	numBoards := (len(lines) - 1)/6

	for i := 0; i < numBoards; i++ {
		g.boards = append(g.boards, newBoard(lines[2+6*i:2+5+1+6*i]))
	}
	return g
}

func (s Solution) Part1(strings []string) (string, error) {
	myGame := newGame(strings)
	score, err := myGame.findBoardScoreForStrategy(winningBoardStrategy)
	if err != nil {
		return "", fmt.Errorf("error finding winning board: %w", err)
	}
	return strconv.Itoa(score), nil
}

func (s Solution) Part2(strings []string) (string, error) {
	myGame := newGame(strings)
	score, err := myGame.findBoardScoreForStrategy(losingBoardStrategy)
	if err != nil {
		return "", fmt.Errorf("error finding losing board: %w", err)
	}
	return strconv.Itoa(score), nil
}
