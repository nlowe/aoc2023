package day5

import (
	"fmt"
	"io"
	"iter"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
	"github.com/nlowe/aoc2023/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	sections, stop := iter.Pull(challenge.Sections(input))
	defer stop()

	seeds := strings.Fields(util.MustPull(sections))[1:]

	seedRanges := make([]span, len(seeds)/2)
	for i := range seedRanges {
		seedRanges[i] = span{
			start: util.MustAtoI(seeds[i*2]),
		}
	}

	converter := makeConverter(sections)

	var locations []span
	for _, seed := range seedRanges {
		locations = append(locations, converter(seed)...)
	}

	slices.SortFunc(locations, span.Less)
	return locations[0].start
}
