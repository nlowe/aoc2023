package day5

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"

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

type span struct {
	start  int
	length int
}

func (s span) contains(v int) bool {
	return v >= s.start && v < (s.start+s.length)
}

// intersect returns the start, overlap, and end spans of other compared against s
//
// # For example, given s..s and o..o
//
// # There are five cases
//
//  1. S entirely contains O
//     ..|s ss s|
//     ....|oo|
//     Which returns |s|, |oo|, |s|
//
//  2. S partially overlaps with O at the start
//     ..|sss s|
//     ......|o o|
//     Which returns |sss|, |o|, ||
//
//  3. S partially overlaps with O at the end
//     ....|s sss|
//     ..|o o|
//     Which returns ||, |o|, |sss|
//
//  4. S is entirely contained within O
//     ....|s|
//     ..|o o o|
//     Which returns ||, |s|, ||
//
//  4. S does not overlap with O
//     ..|sss|
//     ..........|ooo|
//     Which returns |sss|, ||, ||
func (s span) intersect(other span) (span, span, span) {
	return s, s, s
}

func convertVia(mapping map[span]span, v int) int {
	for k, d := range mapping {
		if k.contains(v) {
			return d.start + (v - k.start)
		}
	}

	return v
}

func partA(challenge *challenge.Input) int {
	parts := challenge.Sections()

	seeds := strings.Fields(<-parts)[1:]

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
