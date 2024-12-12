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

func isReportSafe(r []int) bool {
	shouldIncrease := false

	if r[0] < r[1] {
		shouldIncrease = true
	}

	for p1, p2 := 0, 1; p2 < len(r); p1, p2 = p1+1, p2+1 {
		val1 := r[p1]
		val2 := r[p2]

		if val1 == val2 {
			return false
		}

		if shouldIncrease {
			if val2 < val1 {
				return false
			}

			diff := val2 - val1

			if diff > 3 {
				return false
			}
		} else {
			if val2 > val1 {
				return false
			}

			diff := val1 - val2
			if diff > 3 {
				return false
			}
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

// func day2Part2(r io.Reader) (int, error) {
// 	var safeReports int

// 	reports, err := buildReports(r)

// 	if err != nil {
// 		return 0, err
// 	}

// 	for _, report := range reports {
// 		res := isReportSafe(report)
// 		if res.safe {
// 			safeReports++
// 			continue
// 		}

// 		// If still not safe, the report can be considered unsafe
// 	}

// 	return safeReports, nil
// }
