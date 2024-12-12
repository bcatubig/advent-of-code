package day1_test

import (
	"os"
	"testing"

	day1 "github.com/bcatubig/advent-of-code/2024/go/day_1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {

	t.Run("example", func(t *testing.T) {

		testFile, err := os.Open("./testdata/input")
		require.Nil(t, err)
		defer testFile.Close()

		want := 11
		got, err := day1.Part1(testFile)
		require.Nil(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("part1", func(t *testing.T) {

		testFile, err := os.Open("../../inputs/day1")
		require.Nil(t, err)
		defer testFile.Close()

		want := 1873376
		got, err := day1.Part1(testFile)
		require.Nil(t, err)
		assert.Equal(t, want, got)
	})
}

func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		testFile, err := os.Open("./testdata/input")
		require.Nil(t, err)
		defer testFile.Close()

		want := 31
		got, err := day1.Part2(testFile)
		require.Nil(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("part2", func(t *testing.T) {
		testFile, err := os.Open("../../inputs/day1")
		require.Nil(t, err)
		defer testFile.Close()

		want := 18997088
		got, err := day1.Part2(testFile)
		require.Nil(t, err)
		assert.Equal(t, want, got)
	})
}
