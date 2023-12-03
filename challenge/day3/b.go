package day3

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
	"github.com/nlowe/aoc2023/util/tilemap"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	var sum int

	schema := tilemap.FromInput(challenge)
	schema.Walk(func(_ rune, x int, y int) {
		sum += ratio(schema, x, y)
	})

	return sum
}

func ratio(schema *tilemap.Map[rune], x, y int) int {
	r, ok := schema.TileAt(x, y)
	if !ok || r != '*' {
		return 0
	}

	var adjacent []int

	w, h := schema.Size()
	seen := tilemap.Of[bool](w, h)

	schema.WalkAllNeighbors(x, y, func(r rune, cx int, cy int) {
		if candidate, ok := schema.TileAt(cx, cy); !ok || candidate < '0' || candidate > '9' {
			// Outside the map or not a digit
			return
		}

		if seenAlready, _ := seen.TileAt(cx, cy); seenAlready {
			// This digit belongs to a number we've already seen
			return
		}

		v, start, end := extractNumber(schema, cx, cy)
		adjacent = append(adjacent, v)
		for nx := start; nx <= end; nx++ {
			// Mark the digits of this number so we don't double-count them
			seen.SetTile(nx, cy, true)
		}
	})

	// If there aren't exactly two adjacent numbers, this isn't a gear so return a ratio of 0
	if len(adjacent) != 2 {
		return 0
	}

	return adjacent[0] * adjacent[1]
}
