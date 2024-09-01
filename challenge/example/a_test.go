package example

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := strings.NewReader("42")

	result := a(input)

	require.Equal(t, 42, result)
}
