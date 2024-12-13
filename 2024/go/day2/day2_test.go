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
		safe, _ := isReportSafe([]int{7, 6, 4, 2, 1})
		assert.True(t, safe)
	})

	t.Run("case 2 - not safe", func(t *testing.T) {
		safe, badIdx := isReportSafe([]int{1, 2, 7, 8, 9})
		assert.False(t, safe)
		assert.Equal(t, 2, badIdx)
	})

}

func TestIsReportSafeWithDapener(t *testing.T) {
	t.Run("case 1 - safe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{7, 6, 4, 2, 1})
		assert.True(t, got)
	})

	t.Run("case 2 - not safe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{1, 2, 7, 8, 9})
		assert.False(t, got)
	})

	t.Run("case 3 - not safe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{9, 7, 6, 2, 1})
		assert.False(t, got)
	})

	t.Run("case 4 - safe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{1, 3, 2, 4, 5})
		assert.True(t, got)
	})

	t.Run("case 5 - not safe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{8, 6, 4, 4, 1})
		assert.True(t, got)
	})

	t.Run("case 6 - safe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{1, 3, 6, 7, 9})
		assert.True(t, got)
	})

	t.Run("case 7 - unsafe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{9, 7, 9, 8})
		assert.False(t, got)
	})

	t.Run("case 8 - safe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{3, 6, 7, 9, 11, 8})
		assert.True(t, got)
	})

	t.Run("case 9 - safe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{21, 24, 25, 26, 29, 30, 32, 32})
		assert.True(t, got)
	})

	t.Run("case 10 - unsafe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{43, 44, 45, 44, 46, 44})
		assert.False(t, got)
	})

	t.Run("case 11 - unsafe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{73, 75, 73, 74, 75, 75})
		assert.False(t, got)
	})

	t.Run("case 12 - unsafe", func(t *testing.T) {
		got := isReportSafeWithDapener([]int{71, 69, 70, 71, 72, 75})
		assert.True(t, got)
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

func TestDay2Part2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		fData, err := os.Open("./testdata/input")
		require.Nil(t, err)
		defer fData.Close()

		got, err := day2Part2(fData)
		require.Nil(t, err)

		assert.Equal(t, 4, got)
	})

	t.Run("input", func(t *testing.T) {
		fData, err := os.Open("../../inputs/day2")
		require.Nil(t, err)
		defer fData.Close()

		got, err := day2Part2(fData)
		require.Nil(t, err)

		assert.Equal(t, 717, got)
	})
}
