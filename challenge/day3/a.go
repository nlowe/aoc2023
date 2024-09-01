package day3

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2023/challenge"
	"github.com/nlowe/aoc2023/util/tilemap"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {
	var sum int
	schema := tilemap.FromInput(input)

	for r, pos := range schema.Values() {
		// Skip this tile if it's not a number
		if r < '0' || r > '9' {
			continue
		}

		// Skip this tile if it doesn't have adjacent symbols
		if !hasAdjacentSymbol(schema, pos.X, pos.Y) {
			continue
		}

		// Otherwise, we found a "part number"
		v, start, end := extractNumber(schema, pos.X, pos.Y)

		sum += v
		for nx := start; nx <= end; nx++ {
			// Consume this part number by masking it out so we don't double count
			schema.SetTile(nx, pos.Y, '.')
		}
	}

	return sum
}

func hasAdjacentSymbol(schema *tilemap.Map[rune], x, y int) bool {
	for r := range schema.AllNeighbors(x, y) {
		if r != '.' && !(r >= '0' && r <= '9') {
			return true
		}
	}

	return false
}

func extractNumber(schema *tilemap.Map[rune], x, y int) (int, int, int) {
	w, _ := schema.Size()

	// Move x to start of number
	for x > 0 {
		x--

		if r, ok := schema.TileAt(x, y); !ok || r < '0' || r > '9' {
			// Went to far, back up one
			x++
			break
		}
	}

	// save the start position
	sx := x

	// Construct v from consecutive digits
	var v int
	for x < w {
		r, ok := schema.TileAt(x, y)
		if !ok || r < '0' || r > '9' {
			break
		}

		v *= 10
		v += int(r - '0')
		x++
	}

	return v, sx, x - 1
}
