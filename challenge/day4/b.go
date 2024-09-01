package day4

import (
	"fmt"
	"io"
	"slices"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	cards := slices.Collect(challenge.Lines(input))

	// Each card starts with one copy
	counts := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		counts[i] = 1
	}

	for i, card := range cards {
		if have := scratchOff(card); have > 0 {
			// Increase the count of the next <have> cards by the number of copies of this card
			// The problem guarantees this won't ever over-run the list of cards
			for j := i + 1; j <= i+have; j++ {
				counts[j] += counts[i]
			}
		}
	}

	var sum int
	for _, v := range counts {
		sum += v
	}

	return sum
}
