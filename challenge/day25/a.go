package day25

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 25, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	panic("Not implemented!")
}