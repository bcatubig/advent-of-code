package day3

import (
	"bufio"
	"fmt"
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

func parseMuls2(in io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		for i := 0; i < lineLen; i++ {
			ch := string(line[i])
			switch ch {
			case "m":
				// Peek ahead 3
				if i+3 < lineLen {
					if line[i:i+3] == "mul" {
						i += 3
						fmt.Println("We got a mul, boys")
					}
				}
			case "(":
				fmt.Println("left brace")
			case ")":
				fmt.Println("right brace")
			case ",":
				fmt.Println("comma")
			}
		}
	}

	return nil, nil
}

func calculateMulSum(in [][]int) int {
	var result int
	for _, mulSet := range in {
		result += (mulSet[0] * mulSet[1])

	}
	return result
}
