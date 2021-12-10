package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"advent2021/juansc"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"advent2021/joage"
	"advent2021/lib"
)

const (
	dayArg    = "day"
	solverArg = "solver"
)

func run(c *cli.Context) error {
	day := c.Int(dayArg)
	implementer := strings.TrimSpace(c.String(solverArg))

	var solvers lib.DaySolver
	switch implementer {
	case "joage":
		solvers = &joage.Solvers{}
	case "juansc":
		solvers = &juansc.Solvers{}
	default:
		return fmt.Errorf("the implementer %s does not exists", implementer)
	}

	solution, err := solvers.GetSolver(day)
	if err != nil {
		log.Error().Int("day", day).Str("implementer", implementer).Msg("implementer has not attempted the given day")
		return err
	}

	lines, err := lib.ReadLines(fmt.Sprintf("./%s/inputs/day%d.txt", implementer, day))
	if err != nil {
		panic("could not read file")
	}

	start := time.Now()
	solution1, err := solution.Part1(lines)
	if err != nil {
		panic(fmt.Errorf("encountered error running part1: %w", err))
	}
	fmt.Println("solution for day 1:", solution1)
	solution2, err := solution.Part2(lines)
	if err != nil {
		panic(fmt.Errorf("encountered error running part2: %w", err))
	}
	fmt.Println("solution for day 2:", solution2)
	fmt.Println(fmt.Sprintf("Total execution time: %v", time.Since(start)))
	return nil

}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     dayArg,
				Aliases:  []string{"d"},
				Usage:    "Day in the Advent of Code",
				Required: true,
			},
			&cli.StringFlag{
				Name:     solverArg,
				Aliases:  []string{"s"},
				Usage:    "Indicates the person whose solution you want to run. By default it will run the name in the .config file for solver",
				Required: false,
				FilePath: "./.config",
			},
		},
		Action: run,
	}
	if err := app.Run(os.Args); err != nil {
		log.Error().Err(err).Msg("encountered fatal error")
		os.Exit(1)
	}
}
