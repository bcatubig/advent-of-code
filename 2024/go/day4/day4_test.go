package day4

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsePuzzle(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		data, err := os.Open("./testdata/input")
		require.Nil(t, err)
		defer data.Close()

		got, err := parsePuzzle(data)
		require.Nil(t, err)

		for _, row := range got {
			fmt.Println(row)
		}

	})
}
