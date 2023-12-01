package day1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2023/challenge"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

	result := partA(input)

	require.Equal(t, 142, result)
}
