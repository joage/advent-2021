package day12

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/rs/zerolog/log"
)

// https://adventofcode.com/2021/day/12
type Solution struct{}

type cave struct {
	name       string
	small      bool
	neighbours map[string]*cave
}

type caveSystem struct {
	caves map[string]*cave
}

type searchResult struct {
	paths [][]string
}

func buildCaveSystem(lines []string) (caveSystem, error) {
	s := caveSystem{
		caves: make(map[string]*cave),
	}

	for _, l := range lines {
		caves := strings.Split(l, "-")
		c1 := s.add(caves[0])
		c2 := s.add(caves[1])
		c1.neighbours[c2.name] = c2
		c2.neighbours[c1.name] = c1
	}

	return s, nil
}

func (s *caveSystem) add(name string) *cave {
	c, ok := s.caves[name]
	if !ok {
		c = &cave{
			name:       name,
			small:      unicode.IsLower(rune(name[0])),
			neighbours: make(map[string]*cave),
		}
		s.caves[name] = c
	}

	return c
}

func (s *caveSystem) findPaths(node *cave, currentPath []string, sr *searchResult, allowedTwice string) {
	log.Debug().Str("node", node.name).Strs("currentPath", currentPath).Msg("Finding paths")

	var nodeInPath int
	for _, p := range currentPath {
		if p == node.name {
			nodeInPath++
		}
	}

	if node.small && ((nodeInPath > 0 && node.name != allowedTwice) || nodeInPath > 1) {
		log.Debug().Msg("Skipping already visiting small cave")
		return
	}

	currentPath = append(currentPath, node.name)
	if node.name == "end" {
		sr.paths = append(sr.paths, currentPath)
		log.Debug().Msg("Reached the end")
		return
	}

	for _, n := range node.neighbours {
		// Having recursive calls work on new slices so they don't interfere with each other
		nextPath := strings.Split(strings.Join(currentPath, ","), ",")
		s.findPaths(n, nextPath, sr, allowedTwice)
	}
}

func (s *caveSystem) str() string {
	var connections []string
	for _, c := range s.caves {
		for _, n := range c.neighbours {
			connections = append(connections, fmt.Sprintf("%s->%s", c.name, n.name))
		}
	}

	return strings.Join(connections, "\n")
}

func (s Solution) Part1(lines []string) (string, error) {
	cs, err := buildCaveSystem(lines)
	if err != nil {
		return "", err
	}

	log.Debug().Msg("\n" + cs.str())

	sr := &searchResult{}
	cs.findPaths(cs.caves["start"], nil, sr, "")

	for _, p := range sr.paths {
		log.Debug().Msg(strings.Join(p, ","))
	}

	return strconv.Itoa(len(sr.paths)), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	cs, err := buildCaveSystem(lines)
	if err != nil {
		return "", err
	}

	log.Debug().Msg("\n" + cs.str())

	sr := &searchResult{}
	for _, c := range cs.caves {
		if c.name != "start" && c.name != "end" {
			cs.findPaths(cs.caves["start"], nil, sr, c.name)
		}
	}

	uniquePaths := make(map[string][]string)
	for _, p := range sr.paths {
		uniquePaths[strings.Join(p, ",")] = p
	}

	for _, p := range uniquePaths {
		log.Debug().Msg(strings.Join(p, ","))
	}

	return strconv.Itoa(len(uniquePaths)), nil
}
