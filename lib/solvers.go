package lib

// Solutions accepts a slice of strings and can provide an answer for part 1 and part2 of the advent of code. If a
// day is not implemented it should panic instead of returning an error.
type Solutions interface {
	Part1([]string) (string, error)
	Part2([]string) (string, error)
}

type DaySolver interface{
	GetSolver(day int) (Solutions, error)
}

