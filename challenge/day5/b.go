package day5

import (
	"fmt"
	"slices"
	"strings"

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

func partB(challenge *challenge.Input) int {
	parts := challenge.Sections()

	seeds := strings.Fields(<-parts)[1:]

	seedRanges := make([]span, len(seeds)/2)
	for i := range seedRanges {
		seedRanges[i] = span{
			start: util.MustAtoI(seeds[i*2]),
		}
	}

	converter := makeConverter(parts)

	var locations []span
	for _, seed := range seedRanges {
		locations = append(locations, converter(seed)...)
	}

	slices.SortFunc(locations, span.Less)
	return locations[0].start
}
