package day2

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
	"github.com/nlowe/aoc2023/util"
	"github.com/nlowe/aoc2023/util/gmath"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

const (
	totalRed   = 12
	totalGreen = 13
	totalBlue  = 14
)

func partA(input io.Reader) int {
	var sum int

	for game := range challenge.Lines(input) {
		id, _, legal := score(game)

		if legal {
			sum += id
		}
	}

	return sum
}

func score(game string) (int, int, bool) {
	var maxRed, maxGreen, maxBlue int
	legal := true

	parts := strings.Split(game, ":")
	id := util.MustAtoI(strings.TrimPrefix(parts[0], "Game "))
	draws := strings.Split(parts[1], ";")

	for _, draw := range draws {
		colors := strings.Split(draw, ",")

		for _, color := range colors {
			v := strings.Fields(color)
			count := util.MustAtoI(v[0])

			switch v[1] {
			case "red":
				maxRed = gmath.Max(count, maxRed)
				legal = legal && (count <= totalRed)
			case "green":
				maxGreen = gmath.Max(count, maxGreen)
				legal = legal && (count <= totalGreen)
			case "blue":
				maxBlue = gmath.Max(count, maxBlue)
				legal = legal && (count <= totalBlue)
			default:
				panic(fmt.Sprintf("unknown color %s", v[1]))
			}
		}
	}

	return id, maxRed * maxGreen * maxBlue, legal
}
