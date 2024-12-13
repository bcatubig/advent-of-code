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

func isIncrementing(report []int) bool {
	incCount := 0
	decCount := 0

	for i, j := 0, 1; j < len(report); i, j = i+1, j+1 {
		num1 := report[i]
		num2 := report[j]

		if num1 < num2 {
			incCount++
		} else {
			decCount++
		}
	}

	return incCount > decCount
}

func isReportSafe(report []int) (bool, int) {
	isIncrementing := isIncrementing(report)
	for i, j := 0, 1; j < len(report); i, j = i+1, j+1 {
		num1 := report[i]
		num2 := report[j]

		if num1 == num2 {
			return false, i
		}

		// Check the diffs
		if isIncrementing {
			if num2 < num1 {
				return false, j
			}
			diff := num2 - num1
			if diff > 3 {
				return false, j
			}
		} else {
			if num2 > num1 {
				return false, j
			}

			diff := num1 - num2
			if diff > 3 {
				return false, j
			}
		}
	}

	return true, 0
}

func isReportSafeWithDapener(report []int) bool {
	isSafe, badIdx := isReportSafe(report)

	if isSafe {
		return true
	}

	dampenedReport1 := removeItemFromSlice(report, badIdx)
	isSafe, _ = isReportSafe(dampenedReport1)
	if isSafe {
		return true
	}

	if badIdx-1 >= 0 {
		dampenedReport2 := removeItemFromSlice(report, badIdx-1)
		isSafe, _ = isReportSafe(dampenedReport2)
		if isSafe {
			return true
		}
	}

	return false
}

func day2Part1(r io.Reader) (int, error) {
	var safeReports int

	reports, err := buildReports(r)

	if err != nil {
		return 0, err
	}

	for _, report := range reports {
		safe, _ := isReportSafe(report)
		if safe {
			safeReports++
			continue
		}
	}

	return safeReports, nil
}

func day2Part2(r io.Reader) (int, error) {
	var safeReports int

	reports, err := buildReports(r)

	if err != nil {
		return 0, err
	}

	for _, report := range reports {
		if isReportSafeWithDapener(report) {
			safeReports++
		}
	}

	return safeReports, nil
}
