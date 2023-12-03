package day3

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2023/challenge"
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
	input := challenge.FromLiteral(example)

	result := partA(input)

	require.Equal(t, 4361, result)
}
