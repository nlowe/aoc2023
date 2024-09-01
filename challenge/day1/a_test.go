package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := strings.NewReader(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

	result := partA(input)

	require.Equal(t, 142, result)
}
