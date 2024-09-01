package day3

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
	"github.com/nlowe/aoc2023/util/tilemap"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	var sum int

	schema := tilemap.FromInput(input)
	for _, pos := range schema.Values() {
		sum += ratio(schema, pos.X, pos.Y)
	}

	return sum
}

func ratio(schema *tilemap.Map[rune], x, y int) int {
	r, ok := schema.TileAt(x, y)
	if !ok || r != '*' {
		return 0
	}

	adjacent := make([]int, 0, 8)

	w, h := schema.Size()
	seen := tilemap.Of[bool](w, h)

	for _, pos := range schema.AllNeighbors(x, y) {
		if candidate, ok := schema.TileAt(pos.X, pos.Y); !ok || candidate < '0' || candidate > '9' {
			// Outside the map or not a digit
			continue
		}

		if seenAlready, _ := seen.TileAt(pos.X, pos.Y); seenAlready {
			// This digit belongs to a number we've already seen
			continue
		}

		v, start, end := extractNumber(schema, pos.X, pos.Y)
		adjacent = append(adjacent, v)
		for nx := start; nx <= end; nx++ {
			// Mark the digits of this number so we don't double-count them
			seen.SetTile(nx, pos.Y, true)
		}
	}

	// If there aren't exactly two adjacent numbers, this isn't a gear so return a ratio of 0
	if len(adjacent) != 2 {
		return 0
	}

	return adjacent[0] * adjacent[1]
}
