package day5

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2023/challenge"
)

func TestB(t *testing.T) {
	t.Skipf("still working on this problem")
	input := challenge.FromLiteral(example)

	result := partB(input)

	require.Equal(t, 46, result)
}
