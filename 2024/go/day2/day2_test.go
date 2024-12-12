package day2

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildReports(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		fData, err := os.Open("./testdata/input")
		require.Nil(t, err)
		defer fData.Close()

		got, err := buildReports(fData)
		require.Nil(t, err)

		assert.Equal(t, [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		}, got)
	})
}

func TestIsReportSafe(t *testing.T) {
	t.Run("case 1 - safe", func(t *testing.T) {
		got := isReportSafe([]int{7, 6, 4, 2, 1})
		assert.True(t, got.safe)
	})

	t.Run("case 2 - not safe", func(t *testing.T) {
		got := isReportSafe([]int{1, 2, 7, 8, 9})
		assert.False(t, got.safe)
	})

	t.Run("case 3 - not safe", func(t *testing.T) {
		got := isReportSafe([]int{9, 7, 6, 2, 1})
		assert.False(t, got.safe)
	})

	t.Run("case 4 - not safe", func(t *testing.T) {
		got := isReportSafe([]int{1, 3, 2, 4, 5})
		assert.False(t, got.safe)
	})

	t.Run("case 5 - not safe", func(t *testing.T) {
		got := isReportSafe([]int{8, 6, 4, 4, 1})
		assert.False(t, got.safe)
	})

	t.Run("case 6 - safe", func(t *testing.T) {
		got := isReportSafe([]int{1, 3, 6, 7, 9})
		assert.True(t, got.safe)
	})

}

func TestDay2Part1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		fData, err := os.Open("./testdata/input")
		require.Nil(t, err)
		defer fData.Close()

		got, err := day2Part1(fData)
		require.Nil(t, err)

		assert.Equal(t, 2, got)
	})

	t.Run("input", func(t *testing.T) {
		fData, err := os.Open("../../inputs/day2")
		require.Nil(t, err)
		defer fData.Close()

		got, err := day2Part1(fData)
		require.Nil(t, err)

		assert.Equal(t, 686, got)
	})
}

// func TestDay2Part2(t *testing.T) {
// 	t.Run("example", func(t *testing.T) {
// 		fData, err := os.Open("./testdata/input")
// 		require.Nil(t, err)
// 		defer fData.Close()

// 		got, err := day2Part2(fData)
// 		require.Nil(t, err)

// 		assert.Equal(t, 4, got)
// 	})

// 	t.Run("input", func(t *testing.T) {
// 		fData, err := os.Open("../../inputs/day2")
// 		require.Nil(t, err)
// 		defer fData.Close()

// 		got, err := day2Part2(fData)
// 		require.Nil(t, err)

// 		assert.Equal(t, 686, got)
// 	})
// }
