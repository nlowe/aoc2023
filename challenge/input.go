package challenge

import (
	"bufio"
	"io"
	"iter"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nlowe/aoc2023/util"

	"github.com/spf13/viper"
)

// InputFile returns an io.Reader for the file pointed at by the --input flag. If the flag is not specified, it looks
// for a file named "input.txt" in the same package as the caller.
func InputFile() io.Reader {
	path := viper.GetString("input")
	if path == "" {
		_, f, _, ok := runtime.Caller(1)
		if !ok {
			panic("failed to determine input path, provide it with -i instead")
		}

		path = filepath.Join(filepath.Dir(f), "input.txt")
	}

	r, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return r
}

// Lines returns an iter.Seq[string] over all lines in the provided io.Reader.
func Lines(r io.Reader) iter.Seq[string] {
	scanner := bufio.NewScanner(r)

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if err := scanner.Err(); err != nil && err != io.EOF {
				panic(err)
			}

			if !yield(scanner.Text()) {
				return
			}
		}
	}
}

// Sections returns an iter.Seq[string] over all blocks of lines in the provided io.Reader. Blocks of lines have at
// least one extra newline separating them.
func Sections(r io.Reader) iter.Seq[string] {
	scanner := bufio.NewScanner(r)
	var section strings.Builder

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if err := scanner.Err(); err != nil && err != io.EOF {
				panic(err)
			}

			line := scanner.Text()
			section.WriteString(line)

			if line == "" {
				if !yield(strings.TrimSpace(section.String())) {
					return
				}
				section.Reset()
			} else {
				section.WriteRune('\n')
			}
		}

		if section.Len() != 0 {
			yield(section.String())
		}
	}
}

// Ints returns an iter.Seq[int] over all lines in the provided io.Reader, converting each line to an int. This method
// panics if conversion of any line fails.
func Ints(r io.Reader) iter.Seq[int] {
	scanner := bufio.NewScanner(r)
	return func(yield func(int) bool) {
		for scanner.Scan() {
			err := scanner.Err()
			if err != nil && err != io.EOF {
				panic(err)
			}

			if !yield(util.MustAtoI(scanner.Text())) {
				return
			}
		}
	}
}
