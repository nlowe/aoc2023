package day3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestA(t *testing.T) {
	input := strings.NewReader(example)

	result := partA(input)

	require.Equal(t, 4361, result)
}
