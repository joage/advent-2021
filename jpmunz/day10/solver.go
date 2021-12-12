package day10

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/10
type Solution struct{}

type syntaxError struct {
	token string
}

type stack struct {
	items []string
}

func (s *stack) pop() (string, error) {
	if s.empty() {
		return "", fmt.Errorf("attempt to pop from empty stack")
	}

	i := s.peek()
	s.items = s.items[:len(s.items)-1]

	return i, nil
}

func (s *stack) push(i string) {
	s.items = append(s.items, i)
}

func (s *stack) peek() string {
	if s.empty() {
		return ""
	}

	return s.items[len(s.items)-1]
}

func (s *stack) empty() bool {
	return len(s.items) == 0
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("Unexpected string: %s", e.token)
}

type transition struct {
	state string
	next  string
}

func (t *transition) str() string {
	return fmt.Sprintf("%s -> %s", t.state, t.next)
}

var (
	validTransitions = map[transition]bool{
		transition{state: "", next: "{"}: false,
		transition{state: "", next: "("}: false,
		transition{state: "", next: "["}: false,
		transition{state: "", next: "<"}: false,

		transition{state: "{", next: "}"}: true,
		transition{state: "{", next: "{"}: false,
		transition{state: "{", next: "("}: false,
		transition{state: "{", next: "["}: false,
		transition{state: "{", next: "<"}: false,

		transition{state: "(", next: ")"}: true,
		transition{state: "(", next: "{"}: false,
		transition{state: "(", next: "("}: false,
		transition{state: "(", next: "["}: false,
		transition{state: "(", next: "<"}: false,

		transition{state: "[", next: "]"}: true,
		transition{state: "[", next: "{"}: false,
		transition{state: "[", next: "("}: false,
		transition{state: "[", next: "["}: false,
		transition{state: "[", next: "<"}: false,

		transition{state: "<", next: ">"}: true,
		transition{state: "<", next: "{"}: false,
		transition{state: "<", next: "("}: false,
		transition{state: "<", next: "["}: false,
		transition{state: "<", next: "<"}: false,
	}
	errorPoints = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	completionPoints = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	closers = map[string]string{
		"[": "]",
		"(": ")",
		"<": ">",
		"{": "}",
	}
)

func parseLine(tokens []string) ([]string, error) {
	program := stack{}

	for _, t := range tokens {
		attempt := transition{
			state: program.peek(),
			next:  t,
		}

		closedChunk, ok := validTransitions[attempt]
		if !ok {
			log.Debug().Str("line", strings.Join(tokens, "")).Str("failed_transition", attempt.str()).Msg("Syntax error")
			return nil, &syntaxError{token: t}
		}

		if closedChunk {
			program.pop()
		} else {
			program.push(t)
		}
	}

	var completion []string
	for {
		if program.empty() {
			break
		}

		top, err := program.pop()
		if err != nil {
			return nil, err
		}

		closer, ok := closers[top]
		if !ok {
			return nil, fmt.Errorf("no closure found for %s", top)
		}

		completion = append(completion, closer)
	}

	return completion, nil
}

func (s Solution) Part1(lines []string) (string, error) {
	var sum int
	for _, l := range lines {
		_, err := parseLine(strings.Split(l, ""))
		if err != nil {
			se, ok := err.(*syntaxError)
			if ok {
				sum += errorPoints[se.token]
			} else {
				return "", err
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	var linePoints []int
	for _, l := range lines {
		completion, err := parseLine(strings.Split(l, ""))
		if err != nil {
			continue
		}

		var points int
		for _, c := range completion {
			points *= 5
			points += completionPoints[c]
		}

		log.Debug().Str("line", l).Str("completion", strings.Join(completion, "")).Int("points", points).Msg("Scored line")
		linePoints = append(linePoints, points)
	}

	sort.Ints(linePoints)
	middleIndex := (len(linePoints) - 1) / 2
	log.Debug().Ints("line_points", linePoints).Int("middle_index", middleIndex).Msg("Found middle")

	return strconv.Itoa(linePoints[middleIndex]), nil
}
