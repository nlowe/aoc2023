package day1

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	return sum(challenge, extractDigit)
}

func sum(challenge *challenge.Input, extract func(string) (int, bool)) int {
	var answer int

	for line := range challenge.Lines() {
		var v int
		for i := 0; i < len(line); i++ {
			if vv, ok := extract(line[i:]); ok {
				v = vv * 10
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if vv, ok := extract(line[i:]); ok {
				v += vv
				break
			}
		}

		answer += v
	}

	return answer
}

func extractDigit(line string) (int, bool) {
	if line[0] >= '0' && line[0] <= '9' {
		return int(line[0] - '0'), true
	}

	return 0, false
}
