package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func buildReports(r io.Reader) ([][]int, error) {
	var result [][]int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		report := make([]int, len(fields))

		for idx, val := range fields {
			rVal, err := strconv.Atoi(val)

			if err != nil {
				return nil, err
			}
			report[idx] = rVal
		}

		result = append(result, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func removeItemFromSlice(in []int, targetIdx int) []int {
	var result []int

	for idx, num := range in {
		if idx == targetIdx {
			continue
		}
		result = append(result, num)
	}

	return result
}

type reportResult struct {
	safe   bool
	badIdx int
}

func isReportSafe(r []int) *reportResult {
	result := &reportResult{}

	left, right := 0, 1

	isIncreasing := r[left] < r[right]

	for right < len(r) {
		if r[left] == r[right] {
			return result
		}

		if isIncreasing {
			diff := r[right] - r[left]
			if diff > 3 || diff < 0 {
				return result
			}
		} else {
			diff := r[left] - r[right]
			if diff > 3 || diff < 0 {
				return result
			}
		}

		left++
		right++
	}

	result.safe = true

	return result
}

func day2Part1(r io.Reader) (int, error) {
	var safeReports int

	reports, err := buildReports(r)

	if err != nil {
		return 0, err
	}

	for _, report := range reports {
		res := isReportSafe(report)
		if res.safe {
			safeReports++
			continue
		}
	}

	return safeReports, nil
}

func day2Part2(r io.Reader) (int, error) {
	return 0, nil
}
