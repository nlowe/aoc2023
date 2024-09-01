package util

import "iter"

// First returns the first item in the sequence and false if the sequence was empty.
func First[T any](in iter.Seq[T]) (T, bool) {
	next, stop := iter.Pull(in)
	defer stop()

	return next()
}

// MustPull invokes the pull function of the iterator and panics if there are no more items in the sequence.
func MustPull[T any](next func() (T, bool)) T {
	v, ok := next()
	if !ok {
		panic("expected value but pull iter empty")
	}

	return v
}
