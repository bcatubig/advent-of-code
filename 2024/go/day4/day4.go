package day4

import (
	"bufio"
	"io"
	"strings"
)

func parsePuzzle(r io.Reader) ([][]string, error) {
	var result [][]string

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, ch := range line {
			row = append(row, strings.ToLower(string(ch)))
		}
		result = append(result, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
