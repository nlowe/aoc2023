package day5

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/nlowe/aoc2023/util"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func convertManyVia(mapping map[span]span, v span) []span {
	// Given some input span v..v and this mapping
	//   a..a -> b..b
	//   c..c -> d..d
	//
	// Produce output spans mapped following the rules
	//
	// v vv vvvvvv v v vvvvv
	//  |aa aaaaaa|
	//     |bbbbbb b b|
	//              |c c|
	//                |d dd|
	// v|vv|bbbbbb|v v d|vv|

	// Sort the mapping by the start of each span
	spans := maps.Keys(mapping)
	slices.SortFunc(spans, func(i, j span) int {
		return i.start - j.start
	})

	result := make([]span, 0, 3)

	for _, k := range spans {
		start, overlap, end := mapping[k].intersect(v)

		if start.length > 0 {
			result = append(result, start)
		}

		if overlap.length > 0 {
			result = append(result, overlap)
		}

		if end.length > 0 {
			result = append(result, overlap)
		}
	}

	return result
}

func partB(challenge *challenge.Input) int {
	parts := challenge.Sections()

	seeds := strings.Fields(<-parts)[1:]

	seedRanges := make([]span, len(seeds)/2)
	for i := range seedRanges {
		seedRanges[i] = span{
			start: util.MustAtoI(seeds[i*2]),
		}
	}

	seedToSoil := parseMap(<-parts)
	soilToFertilizer := parseMap(<-parts)
	fertilizerToWater := parseMap(<-parts)
	waterToLight := parseMap(<-parts)
	lightToTemp := parseMap(<-parts)
	tempToHumidity := parseMap(<-parts)
	humidityToLocation := parseMap(<-parts)

	locations := make([]int, len(seeds))
	for i, seed := range seeds {
		v := convertVia(seedToSoil, int(util.MustAtoI(seed)))
		v = convertVia(soilToFertilizer, v)
		v = convertVia(fertilizerToWater, v)
		v = convertVia(waterToLight, v)
		v = convertVia(lightToTemp, v)
		v = convertVia(tempToHumidity, v)
		locations[i] = convertVia(humidityToLocation, v)
	}

	slices.Sort(locations)
	return locations[0]
}
