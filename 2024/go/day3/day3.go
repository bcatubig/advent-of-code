package day3

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

const mulRegexp = `mul\((?P<x>\d{1,3}),(?P<y>\d{1,3})\)`

func parseMuls(in io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(in)

	reMul := regexp.MustCompile(mulRegexp)

	var result [][]int

	for scanner.Scan() {
		line := scanner.Text()
		matches := reMul.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			x, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, err
			}

			y, err := strconv.Atoi(match[2])
			if err != nil {
				return nil, err
			}

			result = append(result, []int{x, y})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func calculateMulSum(in [][]int) int {
	var result int
	for _, mulSet := range in {
		result += (mulSet[0] * mulSet[1])

	}
	return result
}
