package day1

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solution(in *os.File) int {
	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		if len(split) != 2 {
			log.Fatalf("expected 2 ints: got: %v", split)
		}
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	var result int
	for i := 0; i < len(leftList); i++ {
		result += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return result
}
