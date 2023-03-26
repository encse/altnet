package slices

import (
	"errors"
	"math/rand"

	"golang.org/x/exp/constraints"
	lib "golang.org/x/exp/slices"
)

// Contains reports whether v is present in s.
func Contains[E comparable](s []E, v E) bool {
	return lib.Contains(s, v)
}

func Any[E comparable](es []E, pred func(e E) bool) bool {
	for _, e := range es {
		if pred(e) {
			return true
		}
	}
	return false
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S ~[]E, E any](s S) S {
	return lib.Clone(s)
}

// Sort sorts a slice of any ordered type in ascending order.
// Sort may fail to sort correctly when sorting slices of floating-point
// numbers containing Not-a-number (NaN) values.
// Use slices.SortFunc(x, func(a, b float64) bool {return a < b || (math.IsNaN(a) && !math.IsNaN(b))})
// instead if the input may contain NaNs.
func Sort[E constraints.Ordered](x []E) {
	lib.Sort(x)
}

// SortFunc sorts the slice x in ascending order as determined by the less function.
// This sort is not guaranteed to be stable.
//
// SortFunc requires that less is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func SortFunc[E any](x []E, less func(a, b E) bool) {
	lib.SortFunc(x, less)
}

func Take[E comparable](s []E, n int) []E {
	if len(s) < n {
		return s
	}
	res := make([]E, n)
	copy(s[:n], res)
	return res
}

func GetOrDefault[A any](items []A, index int) A {
	if index >= len(items) {
		var a A
		return a
	} else {
		return items[index]
	}
}

func Chunk[E any](items []E, chunkSize int) [][]E {
	var chunks [][]E
	for i := 0; i < len(items); i += chunkSize {
		end := i + chunkSize
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[i:end])
	}
	return chunks
}

func Max[A constraints.Ordered](items []A) A {
	if len(items) == 0 {
		var res A
		return res
	} else {
		res := items[0]
		for _, a := range items {
			if a > res {
				res = a
			}
		}
		return res
	}
}

func Map[A any, B any](items []A, f func(a A) B) []B {
	res := make([]B, 0, len(items))
	for _, a := range items {
		res = append(res, f(a))
	}
	return res
}

func Filter[A any](items []A, f func(a A) bool) []A {
	res := make([]A, 0, len(items))
	for _, a := range items {
		if f(a) {
			res = append(res, a)
		}
	}
	return res
}

// Choose returns one item from the slice choosen at random, an error is
// returned if the slice is empty
func Choose[A any](items []A) (A, error) {
	if len(items) > 0 {
		return items[rand.Intn(len(items))], nil
	}

	var zero A
	return zero, errors.New("slice is empty")
}

func ChooseX[A any](items []A) A {
	n, err := Choose(items)
	if err != nil {
		panic(err)
	}
	return n
}
