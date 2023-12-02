package day2

import (
	"fmt"

	"github.com/nlowe/aoc2023/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	var sum int

	for game := range challenge.Lines() {
		_, power, _ := score(game)
		sum += power
	}

	return sum
}
