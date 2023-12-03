package day3

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2023/challenge"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partB(input)

	require.Equal(t, 467835, result)
}
