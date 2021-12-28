package day14

import (
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/14
type Solution struct{}

func parse(lines []string) (string, map[string]string) {
	template := lines[0]
	rules := make(map[string]string)

	for i := 2; i < len(lines); i++ {
		tokens := strings.Split(lines[i], " -> ")
		rules[tokens[0]] = tokens[1]
	}

	return template, rules
}

func applyRules(template string, rules map[string]string) string {
	var updated []string
	tokens := strings.Split(template, "")

	for i := 0; i < len(tokens)-1; i++ {
		pair := tokens[i] + tokens[i+1]
		insert, ok := rules[pair]

		updated = append(updated, tokens[i])
		if ok {
			updated = append(updated, insert)
		}
	}

	updated = append(updated, tokens[len(tokens)-1])

	return strings.Join(updated, "")
}

func (s Solution) Part1(lines []string) (string, error) {
	template, rules := parse(lines)

	for i := 0; i < 10; i++ {
		template = applyRules(template, rules)
		log.Debug().Str("template", template).Int("step", i+1).Msg("Applied rules")
	}

	frequencies := make(map[string]int)
	for _, c := range template {
		frequencies[string(c)] += 1
	}

	var most, least int
	for k, v := range frequencies {
		log.Debug().Str("element", k).Int("frequency", v).Msg("Final frequency")

		if v > most {
			most = v
		}

		if v < least || least == 0 {
			least = v
		}
	}

	return strconv.Itoa(most - least), nil
}

type input struct {
	depth int
	pair  string
}

func computeFrequencies(in input, maxDepth int, rules map[string]string, answers map[input]map[string]int) map[string]int {
	answer, alreadyAnswered := answers[in]
	if alreadyAnswered {
		return answer
	}

	frequencies := make(map[string]int)
	insert, _ := rules[in.pair]
	frequencies[insert]++

	if in.depth < maxDepth {
		lpair := string(in.pair[0]) + insert
		rpair := insert + string(in.pair[1])

		for k, v := range computeFrequencies(input{pair: lpair, depth: in.depth + 1}, maxDepth, rules, answers) {
			frequencies[k] += v
		}

		for k, v := range computeFrequencies(input{pair: rpair, depth: in.depth + 1}, maxDepth, rules, answers) {
			frequencies[k] += v
		}
	}

	answers[in] = frequencies

	return frequencies
}

func (s Solution) Part2(lines []string) (string, error) {
	template, rules := parse(lines)
	tokens := strings.Split(template, "")

	frequencies := make(map[string]int)
	for _, t := range tokens {
		frequencies[t]++
	}

	answers := make(map[input]map[string]int)
	for i := 0; i < len(tokens)-1; i++ {
		for k, v := range computeFrequencies(input{pair: tokens[i] + tokens[i+1], depth: 1}, 40, rules, answers) {
			frequencies[k] += v
		}
	}

	var most, least int
	for k, v := range frequencies {
		log.Debug().Str("element", k).Int("frequency", v).Msg("Final frequency")

		if v > most {
			most = v
		}

		if v < least || least == 0 {
			least = v
		}
	}

	return strconv.Itoa(most - least), nil
}
