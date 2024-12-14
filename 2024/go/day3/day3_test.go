package day3

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseString(t *testing.T) {
	t.Run("example", func(t *testing.T) {

		data, err := os.Open("./testdata/input")
		require.Nil(t, err)
		defer data.Close()

		got, err := parseMuls(data)
		require.Nil(t, err)

		assert.Contains(t, got, []int{2, 4})
		assert.Contains(t, got, []int{5, 5})
		assert.Contains(t, got, []int{11, 8})
		assert.Contains(t, got, []int{8, 5})
	})
}

func TestCalculateMulSum(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		data, err := os.Open("./testdata/input")
		require.Nil(t, err)
		defer data.Close()

		muls, err := parseMuls(data)
		require.Nil(t, err)

		got := calculateMulSum(muls)
		assert.Equal(t, 161, got)
	})

	t.Run("input", func(t *testing.T) {
		data, err := os.Open("../../inputs/day3")
		require.Nil(t, err)
		defer data.Close()

		muls, err := parseMuls(data)
		require.Nil(t, err)

		got := calculateMulSum(muls)
		assert.Equal(t, 188116424, got)
	})
}