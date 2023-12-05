package day5

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"

	"github.com/nlowe/aoc2023/challenge"
	"github.com/nlowe/aoc2023/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func convertVia(mapping map[span]span, v span) []span {
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
	slices.SortFunc(spans, span.Less)

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

func convertManyVia(mapping map[span]span, vs []span) []span {
	result := make([]span, 0, len(vs)*3)
	for _, v := range vs {
		result = append(result, convertVia(mapping, v)...)
	}

	// TODO: We need to collapse neighboring contiguous ranges to prevent memory from exploding

	return result
}

func partA(challenge *challenge.Input) int {
	parts := challenge.Sections()

	seeds := strings.Fields(<-parts)[1:]

	converter := makeConverter(parts)

	var locations []span
	for _, seed := range seeds {
		locations = append(locations, converter(span{start: util.MustAtoI(seed), length: 1})...)
	}

	slices.SortFunc(locations, span.Less)
	return locations[0].start
}

func makeConverter(parts <-chan string) func(span) []span {
	seedToSoil := parseMap(<-parts)
	soilToFertilizer := parseMap(<-parts)
	fertilizerToWater := parseMap(<-parts)
	waterToLight := parseMap(<-parts)
	lightToTemp := parseMap(<-parts)
	tempToHumidity := parseMap(<-parts)
	humidityToLocation := parseMap(<-parts)

	return func(seed span) []span {
		v := convertVia(seedToSoil, seed)
		v = convertManyVia(soilToFertilizer, v)
		v = convertManyVia(fertilizerToWater, v)
		v = convertManyVia(waterToLight, v)
		v = convertManyVia(lightToTemp, v)
		v = convertManyVia(tempToHumidity, v)
		return convertManyVia(humidityToLocation, v)
	}
}

func parseMap(section string) map[span]span {
	result := map[span]span{}

	for _, entry := range strings.Split(section, "\n")[1:] {
		parts := strings.Fields(entry)

		length := util.MustAtoI(parts[2])

		destination := span{start: util.MustAtoI(parts[0]), length: length}
		source := span{start: util.MustAtoI(parts[1]), length: length}

		result[source] = destination
	}

	return result
}
