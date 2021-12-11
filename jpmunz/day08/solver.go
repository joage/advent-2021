package day08

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/8
type Solution struct{}

type digit struct {
	segmentCode string
}

var (
	zero          = digit{segmentCode: "abcefg"}
	one           = digit{segmentCode: "cf"}
	two           = digit{segmentCode: "acdeg"}
	three         = digit{segmentCode: "acdfg"}
	four          = digit{segmentCode: "bcdf"}
	five          = digit{segmentCode: "abdfg"}
	six           = digit{segmentCode: "abdefg"}
	seven         = digit{segmentCode: "acf"}
	eight         = digit{segmentCode: "abcdefg"}
	nine          = digit{segmentCode: "abcdfg"}
	codesToDigits = map[string]string{
		zero.segmentCode:  "0",
		one.segmentCode:   "1",
		two.segmentCode:   "2",
		three.segmentCode: "3",
		four.segmentCode:  "4",
		five.segmentCode:  "5",
		six.segmentCode:   "6",
		seven.segmentCode: "7",
		eight.segmentCode: "8",
		nine.segmentCode:  "9",
	}
)

type display struct {
	possibleArrangements [][]string
	actualArrangement    []string
}

func (s Solution) Part1(lines []string) (string, error) {
	_, values := parseInput(lines)

	var occurrences int
	for _, valueLine := range values {
		for _, v := range valueLine {
			if len(v) == len(one.segmentCode) ||
				len(v) == len(four.segmentCode) ||
				len(v) == len(seven.segmentCode) ||
				len(v) == len(eight.segmentCode) {
				occurrences++
			}
		}
	}

	return strconv.Itoa(occurrences), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	signals, values := parseInput(lines)

	var sum int
	for i := 0; i < len(lines); i++ {
		n, err := decodeToDecimal(signals[i], values[i])
		if err != nil {
			return "", err
		}

		sum += n
	}

	return strconv.Itoa(sum), nil
}

func newDisplay() display {
	d := display{
		possibleArrangements: permutations([]string{"a", "b", "c", "d", "e", "f", "g"}),
	}

	return d
}

func decodeToDigit(arrangement []string, s string) (string, error) {
	inputMapping := strings.Join(arrangement, "")
	outputMapping := []string{"a", "b", "c", "d", "e", "f", "g"}

	decoded := make([]string, len(s))
	for i, c := range s {
		strIndex := strings.Index(inputMapping, string(c))
		decoded[i] = outputMapping[strIndex]
	}

	sort.Strings(decoded)
	code := strings.Join(decoded, "")

	log.Trace().Str("input", s).Str("output", code).Msg("Attempted decode")

	d, ok := codesToDigits[code]
	if !ok {
		return "", fmt.Errorf("No mapping found for code %s", code)
	} else {
		return d, nil
	}
}

func decodeToDecimal(signals, values []string) (int, error) {
	d := newDisplay()

	for _, p := range d.possibleArrangements {
		if tryArrangement(p, signals) {
			d.actualArrangement = p
			break
		}
	}

	if d.actualArrangement == nil {
		return 0, fmt.Errorf("unable to find solution")
	}

	var digits []string
	for _, v := range values {
		n, err := decodeToDigit(d.actualArrangement, v)
		if err != nil {
			return 0, err
		}
		log.Trace().Str("value", v).Str("digit", n).Msg("Converted value to digit")

		digits = append(digits, n)
	}

	return strconv.Atoi(strings.Join(digits, ""))
}

func tryArrangement(arrangement, signals []string) bool {
	for _, s := range signals {
		d, err := decodeToDigit(arrangement, s)
		if err != nil {
			log.Debug().Strs("arrangement", arrangement).Str("signal", s).Msg("Arrangement failed to decode a signal")
			return false
		}

		log.Trace().Strs("arrangement", arrangement).Str("signal", s).Str("digit", d).Msg("Arrangement decoded a signal successfully")
	}

	log.Debug().Strs("arrangement", arrangement).Msg("Arrangement succeeded")
	return true
}

func parseInput(lines []string) ([][]string, [][]string) {
	var signals [][]string
	var values [][]string
	for _, l := range lines {
		inputs := strings.Split(l, " | ")
		signals = append(signals, strings.Split(inputs[0], " "))
		values = append(values, strings.Split(inputs[1], " "))
	}

	return signals, values
}

// https://go.dev/play/p/Ulyo1H2Bii
func permutations(arr []string) [][]string {
	var helper func([]string, int)
	var res [][]string

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
