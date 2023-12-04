package day4

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	var sum int

	for card := range challenge.Lines() {
		if have := scratchOff(card); have > 0 {
			// A winning card starts with a base score of 1, and each matching number we have doubles its score
			// This is easy to calculate by shifting a one by one-less than the number of matches we have
			sum += 1 << (have - 1)
		}
	}

	return sum
}

func scratchOff(card string) int {
	// Isolate the numbers
	parts := strings.Split(card, ":")
	parts = strings.Split(parts[1], "|")

	// Split out the winning numbers and our numbers
	winning := strings.Fields(parts[0])
	ours := strings.Fields(parts[1])

	// And count how many match. These lists will be very small so a simple Contains should be sufficient, if they were
	// larger, we could consider sorting ahead of time and doing a more efficient search since ordering does not matter
	var have int
	for _, n := range ours {
		if slices.Contains(winning, n) {
			have++
		}
	}

	return have
}
