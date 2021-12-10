package day03

import (
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/3
type Solution struct{}

type freq struct {
	one_count  int
	zero_count int
}

type BitCriteria int

const (
	Most BitCriteria = iota
	Least
)

func (s Solution) Part1(lines []string) (string, error) {
	report := getReport(lines)
	frequencies := calculateFrequencies(report)
	gamma := make([]string, len(frequencies))
	epsilon := make([]string, len(frequencies))

	for i, c := range frequencies {
		if c.one_count > c.zero_count {
			gamma[i] = "1"
			epsilon[i] = "0"
		} else {
			gamma[i] = "0"
			epsilon[i] = "1"
		}
	}
	d_gamma := toDecimal(gamma)
	d_epsilon := toDecimal(epsilon)

	log.Info().Int64("gamma", d_gamma).Int64("epsilon", d_epsilon).Msg("Done")

	return strconv.FormatInt(d_gamma*d_epsilon, 10), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	report := getReport(lines)
	frequencies := calculateFrequencies(report)

	ox_rating := findFinal(0, report, frequencies, Most)
	co_rating := findFinal(0, report, frequencies, Least)

	d_ox_rating := toDecimal(ox_rating)
	d_co_rating := toDecimal(co_rating)

	log.Info().Int64("oxygen", d_ox_rating).Int64("co2", d_co_rating).Msg("Done")

	return strconv.FormatInt(d_ox_rating*d_co_rating, 10), nil
}

func findFinal(p int, numbers [][]string, frequencies []freq, criteria BitCriteria) []string {
	filtered := filterBitCriteria(p, numbers, frequencies, criteria)

	if p >= 20 {
		log.Error().Msg("Recursed too far")
		return filtered[0]
	}
	if len(filtered) == 1 {
		return filtered[0]
	} else {
		return findFinal(p+1, filtered, calculateFrequencies(filtered), criteria)
	}
}

func filterBitCriteria(p int, numbers [][]string, frequencies []freq, criteria BitCriteria) [][]string {
	var result [][]string

	for _, n := range numbers {
		switch criteria {
		case Most:
			if (frequencies[p].one_count >= frequencies[p].zero_count && n[p] == "1") ||
				(frequencies[p].zero_count > frequencies[p].one_count && n[p] == "0") {
				result = append(result, n)
			}
		case Least:
			if (frequencies[p].one_count < frequencies[p].zero_count && n[p] == "1") ||
				(frequencies[p].zero_count <= frequencies[p].one_count && n[p] == "0") {
				result = append(result, n)
			}
		}
	}

	return result
}

func toDecimal(s []string) int64 {
	i, err := strconv.ParseInt(strings.Join(s, ""), 2, 64)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to convert to int")
	}

	return i
}

func getReport(lines []string) [][]string {
	var report [][]string
	for _, l := range lines {
		tokens := strings.Split(l, "")
		report = append(report, tokens)
	}

	return report
}

func calculateFrequencies(numbers [][]string) []freq {
	var frequencies []freq
	for _, tokens := range numbers {
		if frequencies == nil {
			frequencies = make([]freq, len(tokens))
		}

		for i, t := range tokens {
			n, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to convert to int")
			}

			switch n {
			case 0:
				frequencies[i].zero_count++
			case 1:
				frequencies[i].one_count++
			}
		}
	}

	return frequencies
}
