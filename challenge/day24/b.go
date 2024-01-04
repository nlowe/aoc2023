package day24

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 24, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	panic("Not implemented!")
}
