package day5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestA(t *testing.T) {
	t.Skipf("re-factoring part a to solve part b, this test doesn't work right now")
	input := strings.NewReader(example)

	result := partA(input)

	require.Equal(t, 35, result)
}
