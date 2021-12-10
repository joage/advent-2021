# Advent of Code 2021 in Go!!!
Coders sharing their solution for Advent of Code 2021

## Running Locally
Run `go mod tidy` to install.

You can run `go run main.go -d <day> -s <solver>` to run the solution for a given day and given solver.
For local development you can add your name to the .config file and the runner will pull your name.
That way you can run `go run main.go -d <day>`.

## Participating
Create a directory with your name, say 'juansc'. Create a directory with the following structure:
```shell
juansc/
  inputs/
    day1.txt
    day2.txt
    ...
  solvers.go
```

In `juansc/solvers.go`, add the following
```go
type Solvers struct{}

func (s *Solvers) GetSolver(day int) (lib.Solutions, error) {
	switch day {
	case 1:
		return mySolution1Struct{}, nil
	case 2:
		return mySolution2Struct{}, nil
	default:
		return nil, fmt.Errorf("no solver exists for the given day")
	}
}
```

Your solution structs should implement the `lib.Solutions` interface, which has the following interface
```go
type Solutions interface {
	Part1([]string) (string, error)
	Part2([]string) (string, error)
}
```

Finally, add your name to the implementer switch statement in `main.go`.