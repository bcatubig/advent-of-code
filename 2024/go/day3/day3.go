package day3

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const mulRegexp = `mul\((?P<x>\d{1,3}),(?P<y>\d{1,3})\)`
const mul2Regexp = `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`

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

func parseMuls2(in io.Reader) ([][]string, error) {
	var result [][]string
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()

		reMul := regexp.MustCompile(mul2Regexp)
		matches := reMul.FindAllStringSubmatch(line, -1)
		result = append(result, matches...)
	}

	return result, nil
}

func filterOps(ops [][]string) ([][]string, [][]string) {
	var result [][]string
	var disabled [][]string

	var dont bool

	for _, op := range ops {
		if op[0] == "do()" {
			dont = false
		} else if op[0] == "don't()" {
			dont = true
		} else if strings.Contains(op[0], "mul(") {
			if dont {
				disabled = append(disabled, op)
				continue
			}
			result = append(result, op)
		}

	}

	return result, disabled
}

func calculateMulSum(in [][]int) int {
	var result int
	for _, mulSet := range in {
		result += (mulSet[0] * mulSet[1])

	}
	return result
}

func calculateMulSum2(in [][]string) int {
	var result int
	for _, mulSet := range in {
		x, _ := strconv.Atoi(mulSet[1])
		y, _ := strconv.Atoi(mulSet[2])
		result += (x * y)

	}
	return result
}
