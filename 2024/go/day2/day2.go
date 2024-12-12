package day2

import (
	"bufio"
	"io"
	"math"
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

func isReportSafe(r []int) bool {
	shouldIncrease, shouldDecrease := false, false

	if r[0] < r[1] {
		shouldIncrease = true
		shouldDecrease = false
	} else {
		shouldDecrease = true
		shouldIncrease = false
	}

	for p1, p2 := 0, 1; p2 < len(r); p1, p2 = p1+1, p2+1 {
		if r[p1] == r[p2] {
			return false
		}

		if shouldIncrease && r[p2] < r[p1] {
			return false
		}

		if shouldDecrease && r[p2] > r[p1] {
			return false
		}

		diff := int(math.Abs(float64(r[p1] - r[p2])))
		if diff > 3 {
			return false
		}
	}

	return true
}

func day2Part1(r io.Reader) (int, error) {
	var safeReports int

	reports, err := buildReports(r)

	if err != nil {
		return 0, err
	}

	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}

	return safeReports, nil
}
