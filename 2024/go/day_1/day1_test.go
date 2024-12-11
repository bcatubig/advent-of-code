package day1_test

import (
	"os"
	"testing"

	day1 "github.com/bcatubig/advent-of-code/2024/go/day_1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {

	t.Run("example", func(t *testing.T) {

		testFile, err := os.Open("./testdata/input")
		require.Nil(t, err)
		defer testFile.Close()

		want := 11
		got := day1.Solution(testFile)
		assert.Equal(t, want, got)
	})

	t.Run("input", func(t *testing.T) {

		testFile, err := os.Open("../../inputs/day1")
		require.Nil(t, err)
		defer testFile.Close()

		want := 1873376
		got := day1.Solution(testFile)
		assert.Equal(t, want, got)
	})
}
