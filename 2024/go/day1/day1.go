package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type BuildListsResult struct {
	Len   int
	Left  []int
	Right []int
}

func BuildLists(in io.Reader) (*BuildListsResult, error) {
	result := &BuildListsResult{
		Left:  make([]int, 0),
		Right: make([]int, 0),
	}

	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		if len(split) != 2 {
			return nil, fmt.Errorf("expected 2 ints: got: %v", split)
		}
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(leftList) != len(rightList) {
		return nil, fmt.Errorf("lists are not equal in size: %v != %v", len(leftList), len(rightList))
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	result.Left = leftList
	result.Right = rightList
	result.Len = len(leftList)

	return result, nil
}

func Part1(in io.Reader) (int, error) {

	lists, err := BuildLists(in)

	if err != nil {
		return 0, err
	}

	var result int
	for i := 0; i < lists.Len; i++ {
		result += int(math.Abs(float64(lists.Left[i] - lists.Right[i])))
	}

	return result, nil
}

func Part2(in io.Reader) (int, error) {
	numRight := make(map[int]int)

	lists, err := BuildLists(in)

	if err != nil {
		return 0, err
	}

	for _, val := range lists.Right {
		numRight[val]++
	}

	var result int
	for _, val := range lists.Left {
		if count, ok := numRight[val]; !ok {
			continue
		} else {
			result += (val * count)
		}
	}

	return result, nil
}
