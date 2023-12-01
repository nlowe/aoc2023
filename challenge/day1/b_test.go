package day1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2023/challenge"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

	result := partB(input)

	require.Equal(t, 281, result)
}
