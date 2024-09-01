package day5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Skipf("still working on this problem")
	input := strings.NewReader(example)

	result := partB(input)

	require.Equal(t, 46, result)
}
