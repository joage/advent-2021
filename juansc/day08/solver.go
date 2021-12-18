package day08

import (
	"sort"
	"strconv"
	"strings"
)

// This solution is mostly algebraic because I didn't like this puzzle from a programming
// perspective.
// The segments have the following frequencies:
// a = 8
// b = 6
// c = 8
// d = 7
// e = 4
// f = 9
// g = 7
// You can identify b,e,f since they have unique frequencies. Digit 1 is the only number
// with two segments, cf. You've identified f, so we know c. Digit 7 has only three segments
// acf, and we can also infer a. At this point we've identified abcef.
// Digits 4 is the only one with 4 digits, and it contains bcdf. We currently do not know d.
// If you perform a diff against 4 using abcef you should get d. By process of elimination we
// have also found g.

type Solution struct{}

func (s Solution) Part1(lines []string) (string, error) {
	total := 0
	for _, line := range lines {
		numbers := strings.Split(line, " | ")[1]
		for _, num := range strings.Split(numbers, " ") {
			switch len(strings.TrimSpace(num)) {
			case 2, 3, 4, 7:
				total += 1
			}
		}
	}
	return strconv.Itoa(total), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		decoder, _ := newDecoder(parts[0])
		currentNum := 0
		for _, n := range strings.Split(parts[1], " ") {
			currentNum = currentNum*10 + decoder.identifyNumber(n)
		}
		total += currentNum
	}
	return strconv.Itoa(total), nil
}

type segment rune

const (
	segA segment = 'a'
	segB segment = 'b'
	segC segment = 'c'
	segD segment = 'd'
	segE segment = 'e'
	segF segment = 'f'
	segG segment = 'g'
)

type Decoder struct {
	encoder map[segment]segment
	decoder map[segment]segment
}

func (d Decoder) encode(s segment) segment {
	out, ok := d.encoder[s]
	if !ok {
		panic("does not exist")
	}
	return out
}

func (d Decoder) encodeString(segments ...segment) string {
	out := ""
	for _, c := range segments {
		out += string(d.encode(c))
	}
	return out
}

func (d Decoder) identifyNumber(segments string) int {
	decodedSegments := make([]rune, len(segments))
	for i, seg := range segments {
		decodedSegments[i] = rune(d.decoder[segment(seg)])
	}
	sort.Slice(decodedSegments, func(i, j int) bool {
		return decodedSegments[i] < decodedSegments[j]
	})
	out := string(decodedSegments)
	switch out {
	case "abcefg":
		return 0
	case "cf":
		return 1
	case "acdeg":
		return 2
	case "acdfg":
		return 3
	case "bcdf":
		return 4
	case "abdfg":
		return 5
	case "abdefg":
		return 6
	case "acf":
		return 7
	case "abcdefg":
		return 8
	case "abcdfg":
		return 9
	default:
		panic("do not know what the segment is " + out)
	}
}

type segmentSet map[segment]bool

func newDecoder(encoded string) (Decoder, error) {
	var oneSegments segmentSet
	var fourSegments segmentSet
	var sevenSegments segmentSet
	var eightSegments segmentSet
	unknownSegments := []segmentSet{}

	segmentFrequencies := map[segment]int{}

	for _, segments := range strings.Split(encoded, " ") {
		switch len(segments) {
		case 2:
			oneSegments = newSet(segments)
		case 3:
			sevenSegments = newSet(segments)
		case 4:
			fourSegments = newSet(segments)
		case 7:
			eightSegments = newSet(segments)
		default:
			unknownSegments = append(unknownSegments, newSet(segments))
		}
		for _, c := range segments {
			val, _ := segmentFrequencies[segment(c)]
			segmentFrequencies[segment(c)] = val + 1
		}
	}
	d := Decoder{
		encoder: map[segment]segment{},
		decoder: map[segment]segment{},
	}
	// Identify segments with unique frequencies
	for s, f := range segmentFrequencies {
		switch f {
		case 6:
			// Segment B is the only one used 6 times
			d.encoder[segB] = s
		case 4:
			// Segment E is the only one used 4 times
			d.encoder[segE] = s
		case 9:
			// Segment F is the only one used 9 times
			d.encoder[segF] = s
		}
	}

	// The c Segment is the segment in the digit 1 that is not F
	cSegments := oneSegments.diff(newSet(string(d.encoder[segF]))).segments()
	if len(cSegments) != 1 {
		panic("failure identifying c segment")
	}
	d.encoder[segC] = cSegments[0]

	// The a segment is the segments in 7 - segments in 1
	aSegments := sevenSegments.diff(oneSegments).segments()
	if len(aSegments) != 1 {
		panic("failure identifying a segment")
	}
	d.encoder[segA] = aSegments[0]

	// The d segment is the difference between the digit 4 segments and abcef.
	dSegments := fourSegments.diff(
		newSet(d.encodeString(segA, segB, segC, segE, segF)),
	).segments()
	if len(dSegments) != 1 {
		panic("error identifying the d segment")
	}
	d.encoder[segD] = dSegments[0]

	// The g segment is the only one left
	gSegments := eightSegments.diff(
		newSet(d.encodeString(segA, segB, segC, segD, segE, segF)),
	).segments()
	if len(dSegments) != 1 {
		panic("failure identifying g segment")
	}
	d.encoder[segG] = gSegments[0]
	for k, v := range d.encoder {
		d.decoder[v] = k
	}
	return d, nil
}

func newSet(letters string) segmentSet {
	m := segmentSet{}
	for _, c := range letters {
		m[segment(c)] = true
	}
	return m
}

func (s segmentSet) diff(other segmentSet) segmentSet {
	d := segmentSet{}
	for seg := range s {
		if _, ok := other[seg]; !ok {
			d[seg] = true
		}
	}
	return d
}

func (s segmentSet) segments() []segment {
	segments := make([]segment, 0, len(s))
	for seg := range s {
		segments = append(segments, seg)
	}
	return segments
}
