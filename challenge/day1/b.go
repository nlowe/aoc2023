package day1

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	return sum(challenge, extractDigitOrWord)
}

func extractDigitOrWord(line string) (int, bool) {
	switch {
	case line[0] >= '0' && line[0] <= '9':
		return int(line[0] - '0'), true
	case strings.HasPrefix(line, "zero"):
		return 0, true
	case strings.HasPrefix(line, "one"):
		return 1, true
	case strings.HasPrefix(line, "two"):
		return 2, true
	case strings.HasPrefix(line, "three"):
		return 3, true
	case strings.HasPrefix(line, "four"):
		return 4, true
	case strings.HasPrefix(line, "five"):
		return 5, true
	case strings.HasPrefix(line, "six"):
		return 6, true
	case strings.HasPrefix(line, "seven"):
		return 7, true
	case strings.HasPrefix(line, "eight"):
		return 8, true
	case strings.HasPrefix(line, "nine"):
		return 9, true
	default:
		return 0, false
	}
}
