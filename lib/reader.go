package lib

import (
	"bufio"
	"os"
)

func ReadLines(file_name string) ([]string, error) {
	file, err := os.Open("depths.in")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, s.Err()
}
