package day5

type span struct {
	start  int
	length int
}

func (s span) Less(other span) int {
	return s.start - other.start
}

// intersect returns the start, overlap, and end spans of other compared against s
//
// # For example, given s..s and o..o
//
// # There are six cases
//
//  1. S entirely contains O
//     ..|s ss s|
//     ....|oo|
//     Which returns |s|, |oo|, |s|
//
//  2. S partially overlaps with O at the start
//     ..|sss s|
//     ......|o o|
//     Which returns |sss|, |o|, ||
//
//  3. S partially overlaps with O at the end
//     ....|s sss|
//     ..|o o|
//     Which returns ||, |o|, |sss|
//
//  4. S is entirely contained within O
//     ....|s|
//     ..|o o o|
//     Which returns ||, |s|, ||
//
//  4. S does not overlap with O (left)
//     ..|sss|
//     ..........|ooo|
//     Which returns |sss|, ||, ||
//
//  5. S does not overlap with O (right)
//     ..........|sss|
//     ..|ooo|
//     Which returns ||, ||, |sss|
func (s span) intersect(other span) (span, span, span) {
	var prefix span
	if s.start < other.start {
		prefix.start = s.start
		// todo: calculate length
	}

	var suffix span
	if (s.start + s.length) > (other.start + other.length) {
		suffix.start = other.start + other.length + 1
		// todo: calculate length
	}

	// todo: calculate overlap

	// s and other do not overlap
	return prefix, s, suffix
}
